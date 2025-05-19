package app

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/cmd/sessions/app/grpcinterceptors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/cmd/sessions/app/tracker"
	"github.com/EugeneTsydenov/chesshub-sessions-service/config"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/interceptor"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Shutdowner interface {
	Shutdown(ctx context.Context) error
}

type App struct {
	RequestTracker *tracker.RequestTracker

	Config *config.Config
	Logger *logrus.Logger

	Database *postgres.Database

	SessionsRepo port.SessionsRepo

	CreateSessionUseCase  usecase.CreateSessionUseCase
	GetSessionByIdUseCase usecase.GetSessionByIdUseCase
	GetSessionsUseCase    usecase.GetSessionsUseCase
	UpdateSessionUseCase  usecase.UpdateSessionUseCase

	SessionController *grpccontroller.SessionController

	GRPCServer *grpc.Server

	shutdownCh  chan struct{}
	shutdowners []Shutdowner
}

func New(cfg *config.Config) *App {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetReportCaller(true)

	return &App{
		RequestTracker: tracker.NewRequestTracker(logger),
		Config:         cfg,
		Logger:         logger,
		shutdownCh:     make(chan struct{}),
	}
}

func (a *App) InitDeps(ctx context.Context) error {
	d, err := postgres.New(ctx, a.Config.Database.DSN())
	a.Logger.Warn("DSN ", a.Config.Database.DSN())
	if err != nil {
		return err
	}

	err = d.Pool().Ping(ctx)
	if err != nil {
		return err
	}

	a.Database = d
	a.RegisterShutdowner(d)

	a.SessionsRepo = repo.NewPostgresSessionRepository(a.Database)
	a.CreateSessionUseCase = usecase.NewCreateSessionUseCase(a.SessionsRepo)
	a.GetSessionByIdUseCase = usecase.NewGetSessionByIdUseCase(a.SessionsRepo)
	a.GetSessionsUseCase = usecase.NewGetSessionsUseCase(a.SessionsRepo)
	a.UpdateSessionUseCase = usecase.NewUpdateSessionUseCase(a.SessionsRepo)
	a.SessionController = grpccontroller.NewSessionController(a.CreateSessionUseCase, a.GetSessionByIdUseCase, a.GetSessionsUseCase, a.UpdateSessionUseCase)

	return nil
}

func (a *App) SetupGRPCServer() {
	a.GRPCServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(grpcinterceptors.RequestTracking(a.RequestTracker, a.Logger), interceptor.ErrorHandlingInterceptor(a.Logger)),
	)
	reflection.Register(a.GRPCServer)
	sessionsproto.RegisterSessionsServiceServer(a.GRPCServer, a.SessionController)
}

func (a *App) Run(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	p := fmt.Sprintf(":%v", a.Config.App.Port)
	listener, err := net.Listen("tcp", p)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		a.Logger.Info("Starting gRPC server", "port", p)
		if err := a.GRPCServer.Serve(listener); err != nil {
			a.Logger.Error("gRPC server error", "error", err)
		}
	}()

	a.Logger.Info("Application started successfully")

	select {
	case sig := <-sigCh:
		a.Logger.Info("Received shutdown signal", "signal", sig)
	case <-a.shutdownCh:
		a.Logger.Info("Shutdown requested programmatically")
	}

	return a.Shutdown(ctx)
}

func (a *App) Shutdown(ctx context.Context) error {
	a.Logger.Info("Starting graceful shutdown")

	a.RequestTracker.SetShuttingDown(true)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	grpcShutdownDone := make(chan struct{})
	go func() {
		a.Logger.Info("Shutting down gRPC server")
		a.GRPCServer.GracefulStop()
		close(grpcShutdownDone)
	}()

	select {
	case <-grpcShutdownDone:
		a.Logger.Info("gRPC server shut down")
	case <-ctx.Done():
		a.Logger.Warn("gRPC server shutdown timed out, forcing stop")
		a.GRPCServer.Stop()
	}

	if err := a.RequestTracker.WaitForCompletion(ctx); err != nil {
		a.Logger.Error("Timed out waiting for requests to complete", "error", err)
	}

	a.Logger.Info("Waiting for active requests to complete")
	for _, shutdowner := range a.shutdowners {
		if err := shutdowner.Shutdown(ctx); err != nil {
			a.Logger.Error("Error shutting down component", "error", err)
		}
	}

	a.Logger.Info("Graceful shutdown completed")

	return nil
}

func (a *App) RegisterShutdowner(s Shutdowner) {
	a.shutdowners = append(a.shutdowners, s)
}
