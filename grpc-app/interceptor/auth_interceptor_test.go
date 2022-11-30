package interceptor

import (
	"context"
	"net"
	"os"
	"testing"

	"go/service1/config"
	pbAuth "go/service1/grpc-app/protos/auth"
	"go/service1/grpc-app/service/auth"
	"go/service1/shared/models/entities"
	aMock "go/service1/shared/usecase/auth/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// global setting
var (
	ctx         context.Context
	grpcAddress string
	any         gomock.Matcher
)

// client setting
var (
	clientConn *grpc.ClientConn
)

// server setting
var (
	ctrl            *gomock.Controller
	mockAuthUseCase *aMock.MockAuthUseCase
	authInterceptor grpc.UnaryServerInterceptor
)

var dataUser = &pbAuth.RequestRegister_Data{
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserPassword: "123456",
}

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Interceptor")
}

var _ = BeforeSuite(func() {
	// initial global setting
	config.LoadEnvFile()
	ctx = context.Background()
	grpcAddress = os.Getenv("GRPC_ADDRESS")
	any = gomock.Any()

	// initial client setting
	clientConn, _ = grpc.DialContext(ctx, grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// initial server setting
	ctrl = gomock.NewController(GinkgoT())
	mockAuthUseCase = aMock.NewMockAuthUseCase(ctrl)

	// initial mock
	mockAuthUseCase.EXPECT().RegisterUser(any, any).Return(
		&entities.ResponRegisterUser{
			UserName:  dataUser.UserName,
			UserEmail: dataUser.UserEmail,
		}, nil,
	)

	// run with goroutin
	runTestGrpcServer()
})

var _ = AfterEach(func() {
	clientConn.Close()
})

var _ = Describe("Interceptor", func() {
	Context("Test Validation Metadata", func() {
		It("SUCCESS", func() {
			kToken := "hello"
			vToken := "world"
			mdToken := metadata.AppendToOutgoingContext(ctx, kToken, vToken)

			callGrpcServiceAuth := pbAuth.NewAuthRPCClient(clientConn)

			result, err := callGrpcServiceAuth.RegisterUser(mdToken, &pbAuth.RequestRegister{Data: dataUser})

			Expect(err).To(BeNil())
			Expect(result).NotTo(BeNil())
		})

		It("FAIL", func() {
			kToken := "hello..."
			vToken := "world"
			mdToken := metadata.AppendToOutgoingContext(ctx, kToken, vToken)

			callGrpcServiceAuth := pbAuth.NewAuthRPCClient(clientConn)

			_, err := callGrpcServiceAuth.RegisterUser(mdToken, &pbAuth.RequestRegister{Data: dataUser})

			Expect(err).NotTo(BeNil())
		})
	})
})

func runTestGrpcServer() {
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(AuthInterceptor(zap.NewExample())),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	// register rpc
	pbAuth.RegisterAuthRPCServer(
		grpcServer,
		auth.NewAuthService(mockAuthUseCase, zap.NewExample()),
	)

	listen, err := net.Listen("tcp", grpcAddress)
	Expect(err).To(BeNil())

	go grpcServer.Serve(listen)
}
