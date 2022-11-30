package auth

import (
	"context"
	"errors"
	"go/service1/config"
	"strings"
	"time"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"go/service1/shared/models/entities"
	"go/service1/shared/utils"

	// protobuff
	pbAuth "go/service1/grpc-app/protos/auth"
	// mock auth repository
	arMock "go/service1/shared/repository/auth/mocks"
)

var (
	ctx                 context.Context
	ctrl                *gomock.Controller
	mockAuthRepository  *arMock.MockAuthRepository
	authUseCase         AuthUseCase
	any                 gomock.Matcher
	userId, userSession string
)

var dataUser = &pbAuth.RequestRegister_Data{
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserPassword: "123456",
}

func TestProductsUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UseCase Auth Suite")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockAuthRepository = arMock.NewMockAuthRepository(ctrl)
	authUseCase = NewAuthUseCaseImpl(mockAuthRepository, zap.NewExample())
	any = gomock.Any()
	userId = uuid.New().String()
	userSession, _ = utils.EncryptSession(userId, 1800)

	// setting mock for register user repository
	mockAuthRepository.EXPECT().RegisterUser(
		ctx,
		&pbAuth.RequestRegister{
			Data: dataUser,
		},
	).Return(
		&entities.ResponRegisterUser{
			UserName:  dataUser.UserName,
			UserEmail: dataUser.UserEmail,
		},
		nil,
	)

	mockAuthRepository.EXPECT().RegisterUser(any, any).Return(nil, errors.New("Failed register user"))

	// setting mock for login user repository
	mockAuthRepository.EXPECT().LoginUser(
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
		},
		nil,
	)

	mockAuthRepository.EXPECT().LoginUser(any, any).Return(nil, errors.New("Failed register user"))
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Checking Auth Usecase", Ordered, func() {
	Context("NewAuthUseCase", func() {
		It("SUCCESS", func() {
			Expect(authUseCase).NotTo(BeNil())
		})
	})

	Context("Register User", func() {
		It("SUCCESS", func() {
			result, err := authUseCase.RegisterUser(
				ctx,
				&pbAuth.RequestRegister{
					Data: dataUser,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.UserEmail).To(Equal(dataUser.UserEmail))
			Expect(result.UserName).To(Equal(dataUser.UserName))
		})

		It("FAIL", func() {
			_, err := authUseCase.RegisterUser(
				ctx,
				&pbAuth.RequestRegister{
					Data: &pbAuth.RequestRegister_Data{
						UserEmail: "",
					},
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authUseCase.RegisterUser(
				c,
				&pbAuth.RequestRegister{
					Data: dataUser,
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Login User", func() {
		It("SUCCESS", func() {
			result, err := authUseCase.LoginUser(
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
			Expect(result.UserId).To(Equal(userId))
			Expect(result.Session).To(Equal(userSession))
			Expect(len(strings.Split(result.Session, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := authUseCase.LoginUser(
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
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authUseCase.LoginUser(
				c,
				&pbAuth.RequestLogin{
					Data: &pbAuth.RequestLogin_Data{
						UserEmail:      dataUser.UserEmail,
						UserPassword:   dataUser.UserPassword,
						UserRememberMe: false,
					},
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
