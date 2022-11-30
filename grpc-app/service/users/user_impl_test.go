package users

import (
	"context"
	"errors"
	"go/service1/shared/utils"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"go/service1/config"
	pbUsers "go/service1/grpc-app/protos/users"

	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
	uMock "go/service1/shared/usecase/users/mocks"
)

var (
	ctx                 context.Context
	ctrl                *gomock.Controller
	mockUsersUseCase    *uMock.MockUsersUseCase
	usersService        UsersService
	any                 gomock.Matcher
	userid, userSession string
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

func TestUsersService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GRPC Service Users")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockUsersUseCase = uMock.NewMockUsersUseCase(ctrl)
	usersService = NewUserServiceImpl(mockUsersUseCase, zap.NewExample())
	any = gomock.Any()
	userid = uuid.New().String()
	userSession, _ = utils.EncryptSession(userid, 60)

	user.UserId = userid
	user.UserSession = userSession
	product.UserId = userid

	// setting for usecase select user
	mockUsersUseCase.EXPECT().SelectUser(ctx, &pbUsers.RequestSelectUser{
		Data: &pbUsers.RequestSelectUser_Data{
			UserId: userid,
		},
	}).Return(user, nil)

	mockUsersUseCase.EXPECT().SelectUser(any, any).Return(nil, errors.New("user not found"))

	// setting for usecase select session user by id
	mockUsersUseCase.EXPECT().SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
		Data: &pbUsers.RequestSelectSessionUserById_Data{
			UserId: userid,
		},
	}).Return(&entities.ResponSelectSessionUserById{
		UserSession: userSession,
		RememberMe:  false,
	}, nil)

	mockUsersUseCase.EXPECT().SelectSessionUserById(any, any).Return(nil, errors.New("user not found"))
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Test Grpc Service Users", Ordered, func() {
	Context("Select User", func() {
		It("SUCCESS", func() {
			result, err := usersService.SelectUser(ctx, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: userid,
				},
			})

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.UserId).To(Equal(userid))
			Expect(result.Data.UserSession).To(Equal(userSession))
			Expect(len(strings.Split(result.Data.UserSession, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := usersService.SelectUser(ctx, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: userid,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersService.SelectUser(c, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: userid,
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Select Session User By Id", func() {
		It("SUCCESS", func() {
			result, err := usersService.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userid,
				},
			})

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.UserSession).To(Equal(userSession))
			Expect(len(strings.Split(result.Data.UserSession, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := usersService.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userid,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersService.SelectSessionUserById(c, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: userid,
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

})
