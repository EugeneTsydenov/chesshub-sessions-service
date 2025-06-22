package app

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EugeneTsydenov/chesshub-sessions-service/cmd/sessions/app/grpcinterceptors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/cmd/sessions/app/tracker"
	"github.com/EugeneTsydenov/chesshub-sessions-service/config"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/interceptor"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/services"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/postgres/repo"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/geoip"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Shutdowner interface {
	Shutdown(ctx context.Context) error
}

type App struct {
	requestTracker *tracker.RequestTracker

	config *config.Config
	logger *logrus.Logger

	database    *postgres.Database
	geoDatabase *geoip.Database

	postgresSessionQueryFactory postgres.SessionQueryFactory

	locator interfaces.GeoIPLocator

	sessionRepo interfaces.SessionRepo

	sessionService interfaces.SessionService

	sessionFilterBuilder sessionfilter.Builder

	startSessionUseCase usecase.StartSession
	stopSessionUseCase  usecase.StopSession
	listSessionsUseCase usecase.ListSessions

	sessionController *grpccontroller.SessionController

	gRPCServer *grpc.Server

	shutdownCh  chan struct{}
	shutdowners []Shutdowner
}

func New(cfg *config.Config) *App {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetReportCaller(true)

	return &App{
		requestTracker: tracker.NewRequestTracker(logger),
		config:         cfg,
		logger:         logger,
		shutdownCh:     make(chan struct{}),
	}
}

func (a *App) InitDeps(ctx context.Context) error {
	err := a.initPgDatabase(ctx)
	if err != nil {
		a.logger.Info(err)
		return err
	}

	err = a.InitGeoDatabase(ctx)
	if err != nil {
		a.logger.Info(err)
		return err
	}

	a.postgresSessionQueryFactory = postgres.NewSessionQueryFactory()

	a.sessionRepo = repo.NewPostgresSessionRepository(a.database, a.postgresSessionQueryFactory)
	a.locator = geoip.NewLocator(a.geoDatabase)

	a.sessionService = services.NewSessionService(a.locator, a.sessionRepo)

	a.sessionFilterBuilder = sessionfilter.NewBuilder()

	a.startSessionUseCase = usecase.NewStartSession(a.sessionService, a.sessionRepo)
	a.stopSessionUseCase = usecase.NewStopSession(a.sessionService, a.sessionRepo)
	a.listSessionsUseCase = usecase.NewListSessions(a.sessionFilterBuilder, a.sessionRepo)

	a.sessionController = grpccontroller.NewSessionController(a.startSessionUseCase, a.stopSessionUseCase, a.listSessionsUseCase)

	return nil
}

func (a *App) initPgDatabase(ctx context.Context) error {
	d, err := postgres.New(ctx, a.config.Database.DSN())
	if err != nil {
		return err
	}

	err = d.Pool().Ping(ctx)
	if err != nil {
		return err
	}

	a.database = d
	a.RegisterShutdowner(d)

	return nil
}

func (a *App) InitGeoDatabase(_ context.Context) error {
	d, err := geoip.New(a.config.GeoIp.DatabasePath)
	if err != nil {
		return err
	}

	a.geoDatabase = d
	a.RegisterShutdowner(d)

	return nil
}

func (a *App) SetupGRPCServer() {
	a.gRPCServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			grpcinterceptors.RequestTracking(a.requestTracker, a.logger),
			interceptor.ErrorHandlingInterceptor(a.logger),
		),
	)
	reflection.Register(a.gRPCServer)
	sessionsproto.RegisterSessionsServiceServer(a.gRPCServer, a.sessionController)
}

func (a *App) Run(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	p := fmt.Sprintf(":%v", a.config.App.Port)
	listener, err := net.Listen("tcp", p)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		a.logger.Info("Starting gRPC server", "port", p)
		if err := a.gRPCServer.Serve(listener); err != nil {
			a.logger.Error("gRPC server error", "error", err)
		}
	}()

	a.logger.Info("Application started successfully")

	select {
	case sig := <-sigCh:
		a.logger.Info("Received shutdown signal", "signal", sig)
	case <-a.shutdownCh:
		a.logger.Info("Shutdown requested programmatically")
	}

	return a.Shutdown(ctx)
}

func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Starting graceful shutdown")

	a.requestTracker.SetShuttingDown(true)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	grpcShutdownDone := make(chan struct{})
	go func() {
		a.logger.Info("Shutting down gRPC server")
		a.gRPCServer.GracefulStop()
		close(grpcShutdownDone)
	}()

	select {
	case <-grpcShutdownDone:
		a.logger.Info("gRPC server shut down")
	case <-ctx.Done():
		a.logger.Warn("gRPC server shutdown timed out, forcing stop")
		a.gRPCServer.Stop()
	}

	if err := a.requestTracker.WaitForCompletion(ctx); err != nil {
		a.logger.Error("Timed out waiting for requests to complete", "error", err)
	}

	a.logger.Info("Waiting for active requests to complete")
	for _, shutdowner := range a.shutdowners {
		if err := shutdowner.Shutdown(ctx); err != nil {
			a.logger.Error("Error shutting down component", "error", err)
		}
	}

	a.logger.Info("Graceful shutdown completed")

	return nil
}

func (a *App) RegisterShutdowner(s Shutdowner) {
	a.shutdowners = append(a.shutdowners, s)
}
