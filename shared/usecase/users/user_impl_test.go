package users

import (
	"context"
	"errors"
	"go/service1/config"
	pbUsers "go/service1/grpc-app/protos/users"
	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
	urMock "go/service1/shared/repository/users/mocks"
	"go/service1/shared/utils"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var (
	ctx                 context.Context
	ctrl                *gomock.Controller
	mockUsersRepository *urMock.MockUserRepository
	usersUseCase        UsersUseCase
	any                 gomock.Matcher
	userId, userSession string
)

var product = &table.Product{
	UserId:       "",
	ProductId:    uuid.New().String(),
	ProductName:  "Kue Basi",
	CategoryName: "electronic",
	ProductInfo:  "Kue Basi tapi enak",
	ProductPrice: 100000,
	ProductSell:  80000,
	ProductStock: 20,
	ProductImage: "example.com/images/userid/productid/images1.jpg",
}

var user = &table.User{
	UserId:       "",
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserImage:    "example.com/images/userid/profile/images1.jpg",
	UserSession:  "",
	UserStatusId: 1,
	Products: []table.Product{
		*product,
	},
	LastUpdate: time.Now().Format("31-09-2000"),
}

func TestUsersUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UseCase Products Suite")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockUsersRepository = urMock.NewMockUserRepository(ctrl)
	usersUseCase = NewUsersUseCaseImpl(mockUsersRepository, zap.NewExample())
	any = gomock.Any()
	userId = uuid.New().String()
	userSession, _ = utils.EncryptSession(userId, 60)

	user.UserId = userId
	user.UserSession = userSession
	product.UserId = userId

	// setting mock for select user repository
	mockUsersRepository.EXPECT().SelectUser(
		ctx,
		&pbUsers.RequestSelectUser{
			Data: &pbUsers.RequestSelectUser_Data{
				UserId: userId,
			},
		},
	).Return(user, nil)

	mockUsersRepository.EXPECT().SelectUser(any, any).Return(nil, errors.New(("user not found")))

	// setting mock for repository select session user by id
	mockUsersRepository.EXPECT().SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
		Data: &pbUsers.RequestSelectSessionUserById_Data{
			UserId: userId,
		},
	}).Return(&entities.ResponSelectSessionUserById{
		UserSession: userSession,
		RememberMe:  false,
	}, nil)

	mockUsersRepository.EXPECT().SelectSessionUserById(any, any).Return(nil, errors.New("user not found"))
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Checking Products Usecase", Ordered, func() {
	Context("New Users Usecase", func() {
		It("SUCCESS", func() {
			Expect(usersUseCase).NotTo(BeNil())
		})
	})

	Context("Select User", func() {
		It("SUCCESS", func() {
			result, err := usersUseCase.SelectUser(
				ctx,
				&pbUsers.RequestSelectUser{
					Data: &pbUsers.RequestSelectUser_Data{
						UserId: userId,
					},
				},
			)

			Expect(err).To(BeNil())
			Expect(result.UserSession).To(Equal(userSession))
			Expect(len(strings.Split(result.UserSession, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := usersUseCase.SelectUser(
				ctx,
				&pbUsers.RequestSelectUser{
					Data: &pbUsers.RequestSelectUser_Data{
						UserId: userId,
					},
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersUseCase.SelectUser(
				c,
				&pbUsers.RequestSelectUser{
					Data: &pbUsers.RequestSelectUser_Data{
						UserId: userId,
					},
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Select Session User By Id", func() {
		It("SUCCESS", func() {
			result, err := usersUseCase.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userId,
				},
			})

			Expect(err).To(BeNil())
			Expect(result.UserSession).To(Equal(userSession))
			Expect(len(strings.Split(result.UserSession, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := usersUseCase.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userId,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersUseCase.SelectSessionUserById(c, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userId,
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
