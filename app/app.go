package app

//
//import (
//	"fmt"
//	"net"
//	"os"
//
//	"github.com/EugeneTsydenov/chesshub-user-service/config"
//	"github.com/EugeneTsydenov/chesshub-user-service/internal/api/grpc/proto/user"
//	"github.com/EugeneTsydenov/chesshub-user-service/internal/api/grpc/service"
//	"github.com/EugeneTsydenov/chesshub-user-service/internal/app/usecase"
//	"github.com/EugeneTsydenov/chesshub-user-service/internal/infra/data"
//	"github.com/EugeneTsydenov/chesshub-user-service/internal/infra/data/repo"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"google.golang.org/grpc/reflection"
//)
//
//type App struct {
//	grpcServer      *grpc.Server
//	cfg             *config.Config
//	dbPool          data.DbPool
//	userService     *service.UserService
//	registerUseCase usecase.RegisterExecutor
//	userRepo        repo.UserRepo
//}
//
//func New() *App {
//	return &App{}
//}
//
//func (a *App) InitDeps() error {
//	if err := a.initConfig(); err != nil {
//		return fmt.Errorf("failed to initialize config: %w", err)
//	}
//
//	if err := a.initDB(); err != nil {
//		return fmt.Errorf("failed to initialize database: %w", err)
//	}
//
//	if err := a.initRepos(); err != nil {
//		return fmt.Errorf("failed to initialize repositories: %w", err)
//	}
//
//	if err := a.initUseCases(); err != nil {
//		return fmt.Errorf("failed to initialize use cases: %w", err)
//	}
//
//	if err := a.initGrpcServices(); err != nil {
//		return fmt.Errorf("failed to initialize gRPC services: %w", err)
//	}
//
//	a.initGrpcServer()
//
//	return nil
//}
//
//func (a *App) initConfig() error {
//	env := os.Getenv("APP_ENV")
//	cfgPath := os.Getenv("CONFIG_PATH")
//
//	cfg, err := config.Load(env, cfgPath)
//	if err != nil {
//		return err
//	}
//
//	a.cfg = cfg
//	return nil
//}
//
//func (a *App) initDB() error {
//	if a.cfg == nil {
//		return fmt.Errorf("config must be initialized before database")
//	}
//
//	p, err := data.NewDbPool(a.cfg.Database.DSN())
//	if err != nil {
//		return err
//	}
//
//	a.dbPool = p
//	return nil
//}
//
//func (a *App) initRepos() error {
//	if a.dbPool == nil {
//		return fmt.Errorf("database must be initialized before repositories")
//	}
//
//	a.userRepo = repo.NewUserRepo(a.dbPool)
//	return nil
//}
//
//func (a *App) initUseCases() error {
//	if a.userRepo == nil {
//		return fmt.Errorf("user repo must be initialized before use cases")
//	}
//
//	a.registerUseCase = usecase.NewRegisterUseCase(a.userRepo)
//	return nil
//}
//
//func (a *App) initGrpcServices() error {
//	if a.registerUseCase == nil {
//		return fmt.Errorf("use cases must be initialized before gRPC services")
//	}
//
//	a.userService = service.NewUserService(a.registerUseCase)
//	return nil
//}
//
//func (a *App) initGrpcServer() {
//	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
//	reflection.Register(a.grpcServer)
//	user.RegisterUserServiceServer(a.grpcServer, a.userService)
//}
//
//func (a *App) Start() error {
//	if a.cfg == nil || a.grpcServer == nil || a.userService == nil {
//		return fmt.Errorf("app dependencies not properly initialized, please call InitDeps first")
//	}
//
//	port := fmt.Sprintf(":%s", a.cfg.App.Port)
//	l, err := net.Listen("tcp", port)
//	if err != nil {
//		return fmt.Errorf("listen failed on starting grpc server: %w", err)
//	}
//
//	fmt.Printf("gRPC server started on port %s\n", a.cfg.App.Port)
//
//	err = a.grpcServer.Serve(l)
//	if err != nil {
//		return fmt.Errorf("grpc server serve with apperrors: %w", err)
//	}
//
//	return nil
//}
//
//// Stop gracefully stops the application
//func (a *App) Stop() {
//	if a.grpcServer != nil {
//		a.grpcServer.GracefulStop()
//	}
//
//	if a.dbPool != nil {
//		// Close DB connection if needed
//		// a.db.Close()
//	}
//}
