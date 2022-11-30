package listener

import (
	// injector
	"go/service1/common"
	authInjector "go/service1/grpc-app/injector/auth"
	productInjector "go/service1/grpc-app/injector/products"
	userInjector "go/service1/grpc-app/injector/users"
	"go/service1/grpc-app/interceptor"

	// pbService
	pbAuth "go/service1/grpc-app/protos/auth"
	pbProducts "go/service1/grpc-app/protos/products"
	pbUsers "go/service1/grpc-app/protos/users"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"go/service1/config"
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

type Listener interface {
	Run()
}

type listernerImpl struct {
	logger            *zap.Logger
	dbConn            *gorm.DB
	grpcServerOptions []grpc.ServerOption
	grpcServer        *grpc.Server
	grpcHealthServer  *health.Server
	address           string
	done              chan os.Signal
}

func NewListenerImpl() Listener {
	return &listernerImpl{}
}

func (l *listernerImpl) InitMain() {
	// load env file
	config.LoadEnvFile()

	// create logger file
	l.logger = common.BuildLogger()

	// running database
	l.dbConn = config.ConnectDB()

	// passing address
	l.address = os.Getenv("GRPC_ADDRESS")

	if l.address == "" {
		l.address = "127.0.0.1:12345"
	}

	l.done = make(chan os.Signal, 1)

	l.grpcServerOptions = []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.AuthInterceptor(l.logger)),
	}

	l.grpcServer = grpc.NewServer(l.grpcServerOptions...)
	l.grpcHealthServer = health.NewServer()
}

func (l *listernerImpl) RegisterService() {
	// initial service
	authService, _ := authInjector.InitializeAuthService(l.dbConn, l.logger)
	usersService, _ := userInjector.InitializeUsersService(l.dbConn, l.logger)
	productsService, _ := productInjector.InitializeProductsService(l.dbConn, l.logger)
	// imagesService, _ := imagesInjector.InitializeUsersService()

	// register service
	pbAuth.RegisterAuthRPCServer(l.grpcServer, authService)
	pbUsers.RegisterUserRPCServer(l.grpcServer, usersService)
	pbProducts.RegisterProductRPCServer(l.grpcServer, productsService)

	// health service
	l.grpcHealthServer.SetServingStatus(pbUsers.UserRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	l.grpcHealthServer.SetServingStatus(pbAuth.AuthRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	l.grpcHealthServer.SetServingStatus(pbProducts.ProductRPC_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	healthpb.RegisterHealthServer(l.grpcServer, l.grpcHealthServer)
}

func (l *listernerImpl) Run() {
	l.InitMain()

	listen, err := net.Listen("tcp", l.address)

	if err != nil {
		log.Fatalf("failed to listen grpc server, err: %v", err.Error())
	}

	reflection.Register(l.grpcServer)
	l.RegisterService()

	signal.Notify(l.done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = l.grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Print("Grpc service 1 running at " + l.address)
	l.logger.Info("Grpc service 1 running at " + l.address)

	<-l.done

	l.logger.Info("Service is going to stop")
	l.grpcServer.Stop()
	l.logger.Info("Service exited properly")
}
