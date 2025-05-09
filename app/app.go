package app

import (
	"fmt"
	"net"
	"os"

	"github.com/EugeneTsydenov/chesshub-sessions-service/config"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data/repo"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/generated/sessions"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer           *grpc.Server
	cfg                  *config.Config
	dbPool               data.DbPool
	sessionsRepo         repo.SessionsRepo
	createSessionUseCase usecase.CreateSessionUseCase
	sessionsService      *service.SessionsService
}

func New() *App {
	return &App{}
}

func (a *App) InitDeps() error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	if err := a.initDB(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	if err := a.initRepos(); err != nil {
		return fmt.Errorf("failed to initialize repositories: %w", err)
	}

	if err := a.initUseCases(); err != nil {
		return fmt.Errorf("failed to initialize use cases: %w", err)
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

func (a *App) initDB() error {
	if a.cfg == nil {
		return fmt.Errorf("config must be initialized before database")
	}

	p, err := data.NewDbPool(a.cfg.Database.DSN())
	if err != nil {
		return err
	}

	a.dbPool = p
	return nil
}

func (a *App) initRepos() error {
	if a.dbPool == nil {
		return fmt.Errorf("database must be initialized before repositories")
	}

	a.sessionsRepo = repo.NewSessionsRepo(a.dbPool)
	return nil
}

func (a *App) initUseCases() error {
	if a.sessionsRepo == nil {
		return fmt.Errorf("sessions repo must be initialized before use cases")
	}

	a.createSessionUseCase = usecase.NewCreateSessionUseCase(a.sessionsRepo)
	return nil
}

func (a *App) initGrpcServices() error {
	if a.createSessionUseCase == nil {
		return fmt.Errorf("use cases must be initialized before gRPC services")
	}

	a.sessionsService = service.NewSessionsService(a.createSessionUseCase)
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
