package listener

import (
	// injector
	authInjector "go/service1/grpc-app/service/injector/auth"
	productInjector "go/service1/grpc-app/service/injector/products"
	userInjector "go/service1/grpc-app/service/injector/users"

	// pbService
	pbAuth "go/service1/grpc-app/protos/auth"
	pbProducts "go/service1/grpc-app/protos/products"
	pbUsers "go/service1/grpc-app/protos/users"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"go/service1/common"
	"go/service1/config"
	"go/service1/grpc-app/interceptor"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

var (
	logger            *zap.Logger
	dbConn            *gorm.DB
	grpcServerOptions []grpc.ServerOption
	grpcServer        *grpc.Server
	grpcHealthServer  *health.Server
	address           string
	done              chan os.Signal
)

func initMain() {
	// load env file
	config.LoadEnvFile()

	// create logger file
	logger = common.BuildLogger()

	// running database
	dbConn = config.ConnectDB()

	// passing address
	address = os.Getenv("GRPC_ADDRESS")

	done = make(chan os.Signal, 1)

	if address == "" {
		address = "127.0.0.1:12345"
	}

	grpcServerOptions = []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.AuthInterceptor(logger)),
	}

	grpcServer = grpc.NewServer(grpcServerOptions...)
	grpcHealthServer = health.NewServer()
}

func registerService(s *grpc.Server, h *health.Server, db *gorm.DB, log *zap.Logger) {
	// initial service
	authService, _ := authInjector.InitializeAuthService(db, log)
	usersService, _ := userInjector.InitializeUsersService(db, log)
	productsService, _ := productInjector.InitializeProductsService(db, log)
	// imagesService, _ := imagesInjector.InitializeUsersService()

	// register service
	pbAuth.RegisterAuthRPCServer(s, authService)
	pbUsers.RegisterUserRPCServer(s, usersService)
	pbProducts.RegisterProductRPCServer(s, productsService)

	// health service
	h.SetServingStatus(pbUsers.UserRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	h.SetServingStatus(pbAuth.AuthRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	h.SetServingStatus(pbProducts.ProductRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	healthpb.RegisterHealthServer(s, h)
}

func Run() {
	initMain()

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen grpc server, err: %v", err.Error())
	}

	reflection.Register(grpcServer)

	registerService(grpcServer, grpcHealthServer, dbConn, logger)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Print("Grpc service 1 running at " + address)
	logger.Info("Grpc service 1 running at " + address)

	<-done

	logger.Info("Service is going to stop")
	grpcServer.Stop()
	logger.Info("Service exited properly")
}
