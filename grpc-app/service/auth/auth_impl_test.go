package auth

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"go/service1/config"
	"go/service1/shared/utils"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"go/service1/shared/models/entities"

	// protobuff auth
	pbAuth "go/service1/grpc-app/protos/auth"

	// mock usecase auth
	aMock "go/service1/shared/usecase/auth/mocks"
)

var dataUser = &pbAuth.RequestRegister_Data{
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserPassword: "123456",
}

var (
	ctx                 context.Context
	ctrl                *gomock.Controller
	mockAuthUseCase     *aMock.MockAuthUseCase
	authService         AuthService
	any                 gomock.Matcher
	userId, userSession string
)

func TestAuthService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GRPC Products Service")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockAuthUseCase = aMock.NewMockAuthUseCase(ctrl)
	authService = NewAuthService(mockAuthUseCase, zap.NewExample())
	any = gomock.Any()
	userId = uuid.New().String()
	userSession, _ = utils.EncryptSession(userId, 60)

	// setting mock auth usecase register
	mockAuthUseCase.EXPECT().RegisterUser(
		ctx,
		&pbAuth.RequestRegister{
			Data: dataUser,
		},
	).Return(
		&entities.ResponRegisterUser{
			UserName:  dataUser.UserName,
			UserEmail: dataUser.UserEmail,
		}, nil,
	)

	mockAuthUseCase.EXPECT().RegisterUser(any, any).Return(nil, errors.New("failed register user"))

	// setting mock auth usecase login
	mockAuthUseCase.EXPECT().LoginUser(
		ctx,
		&pbAuth.RequestLogin{
			Data: &pbAuth.RequestLogin_Data{
				UserEmail:      dataUser.UserEmail,
				UserPassword:   dataUser.UserPassword,
				UserRememberMe: false,
			},
		},
	).Return(
		&entities.ResponLoginUser{
			UserId:  userId,
			Session: userSession,
		}, nil,
	)

	mockAuthUseCase.EXPECT().LoginUser(any, any).Return(nil, errors.New("failed login user"))
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Test Grpc Service Products", Ordered, func() {
	Context("Register User", func() {
		It("SUCCESS", func() {
			result, err := authService.RegisterUser(ctx, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).To(BeNil())
			Expect(result.Data.UserEmail).To(Equal(dataUser.UserEmail))
			Expect(result.Data.UserName).To(Equal(dataUser.UserName))
		})

		It("FAIL", func() {
			_, err := authService.RegisterUser(ctx, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authService.RegisterUser(ctx, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(ctx.Err()))
		})
	})

	Context("Login User", func() {
		It("SUCCESS", func() {
			result, err := authService.LoginUser(
				ctx,
				&pbAuth.RequestLogin{
					Data: &pbAuth.RequestLogin_Data{
						UserEmail:      dataUser.UserEmail,
						UserPassword:   dataUser.UserPassword,
						UserRememberMe: false,
					},
				},
			)

			Expect(err).To(BeNil())
			Expect(result.Data.UserId).To(Equal(userId))
			Expect(result.Data.Session).To(Equal(userSession))
			Expect(len(strings.Split(result.Data.Session, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := authService.LoginUser(
				ctx,
				&pbAuth.RequestLogin{
					Data: &pbAuth.RequestLogin_Data{
						UserEmail:      dataUser.UserEmail,
						UserPassword:   dataUser.UserPassword,
						UserRememberMe: false,
					},
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authService.LoginUser(
				ctx,
				&pbAuth.RequestLogin{
					Data: &pbAuth.RequestLogin_Data{
						UserEmail:      dataUser.UserEmail,
						UserPassword:   dataUser.UserPassword,
						UserRememberMe: false,
					},
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(ctx.Err()))
		})
	})
})
