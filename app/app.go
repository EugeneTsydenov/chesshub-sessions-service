package app

import (
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/config"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/generated/sessions"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

type App struct {
	grpcServer      *grpc.Server
	cfg             *config.Config
	sessionsService *service.SessionsService
}

func New() *App {
	return &App{}
}

func (a *App) InitDeps() error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	if err := a.initGrpcServices(); err != nil {
		return fmt.Errorf("failed to initialize gRPC services: %w", err)
	}

	a.initGrpcServer()

	return nil
}

func (a *App) initConfig() error {
	env := os.Getenv("APP_ENV")
	cfgPath := os.Getenv("CONFIG_PATH")

	cfg, err := config.Load(env, cfgPath)
	if err != nil {
		return err
	}

	a.cfg = cfg
	return nil
}

func (a *App) initGrpcServices() error {
	a.sessionsService = service.NewSessionsService()
	return nil
}

func (a *App) initGrpcServer() {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	sessions.RegisterSessionsServiceServer(a.grpcServer, a.sessionsService)
}

func (a *App) Start() error {
	if a.cfg == nil || a.grpcServer == nil || a.sessionsService == nil {
		return fmt.Errorf("app dependencies not properly initialized, please call InitDeps first")
	}

	port := fmt.Sprintf(":%s", a.cfg.App.Port)
	l, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("listen failed on starting grpc server: %w", err)
	}

	fmt.Printf("gRPC server started on port %s\n", a.cfg.App.Port)

	err = a.grpcServer.Serve(l)
	if err != nil {
		return fmt.Errorf("grpc server serve with apperrors: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	if a.grpcServer != nil {
		a.grpcServer.GracefulStop()
	}
}
