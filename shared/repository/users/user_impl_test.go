package users

import (
	"context"
	"go/service1/common"
	"go/service1/config"
	pbAuth "go/service1/grpc-app/protos/auth"
	pbProducts "go/service1/grpc-app/protos/products"
	pbUsers "go/service1/grpc-app/protos/users"
	"go/service1/shared/models/entities/table"
	"go/service1/shared/repository/auth"
	"go/service1/shared/repository/products"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var authRepo auth.AuthRepository
var usersRepo UserRepository
var productsRepo products.ProductRepository
var dbConn *gorm.DB
var logger *zap.Logger
var ctx context.Context

var dataUser = &pbAuth.RequestRegister_Data{
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserPassword: "123456",
}

var dataProduct = &pbProducts.Product{
	UserId:          "",
	ProductId:       "",
	ProductName:     "Kue Basi",
	ProductCategory: "electronic",
	ProductInfo:     "Kue Basi tapi enak",
	ProductPrice:    100000,
	ProductSell:     80000,
	ProductStock:    20,
	ProductImage:    "example.com/images/userid/productid/images1.jpg",
}

func TestUsersRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Users Repo Suite")
}

var _ = BeforeSuite(func() {
	ctx = context.Background()
	config.LoadEnvFile()
	dbConn = config.ConnectDB()
	logger = common.BuildLogger()
	authRepo = auth.NewAuthRepositoryImpl(dbConn, zap.NewExample())
	usersRepo = NewUsersRepositoryImpl(dbConn, zap.NewExample())
	productsRepo = products.NewProductsRepositoryImpl(dbConn, zap.NewExample())
})

var _ = AfterSuite(func() {
	// delete user register
	dbConn.Where("user_email = ?", dataUser.UserEmail).Delete(&table.User{})
})

var _ = Describe("Checking Users Repository", Ordered, func() {
	Context("Select User", func() {
		It("SUCCESS", func() {
			registerUser, err := authRepo.RegisterUser(ctx, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).To(BeNil())
			Expect(registerUser.UserEmail).To(Equal(dataUser.UserEmail))
			Expect(registerUser.UserName).To(Equal(dataUser.UserName))

			// check if user was stored in db or not
			loginUser, err := authRepo.LoginUser(ctx, &pbAuth.RequestLogin{
				Data: &pbAuth.RequestLogin_Data{
					UserEmail:      dataUser.UserEmail,
					UserPassword:   dataUser.UserPassword,
					UserRememberMe: false,
				},
			})

			Expect(err).To(BeNil())
			Expect(loginUser.UserId).ToNot(Equal(""))
			Expect(len(strings.Split(loginUser.Session, "."))).To(Equal(3))

			// passing user id to data product
			dataProduct.UserId = loginUser.UserId

			// insert data product with current user
			insertProduct, err := productsRepo.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).To(BeNil())
			Expect(insertProduct).NotTo(BeNil())

			selectUser, err := usersRepo.SelectUser(ctx, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: dataProduct.UserId,
				},
			})

			Expect(err).To(BeNil())
			Expect(selectUser.UserId).To(Equal(dataProduct.UserId))
			Expect(selectUser.UserStatusId).To(Equal(uint32(1)))
			Expect(selectUser.UserEmail).To(Equal(dataUser.UserEmail))
			Expect(selectUser.UserName).To(Equal(dataUser.UserName))
			Expect(selectUser.UserPassword).To(Equal(dataUser.UserPassword))
			Expect(len(selectUser.Products)).ToNot(Equal(0))
			Expect(selectUser.Products[0].UserId).To(Equal(dataProduct.UserId))
			Expect(selectUser.Products[0].ProductName).To(Equal(dataProduct.ProductName))
			Expect(selectUser.Products[0].ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(selectUser.Products[0].ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(selectUser.Products[0].ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(selectUser.Products[0].ProductStock).To(Equal(dataProduct.ProductStock))

		})

		It("FAIL", func() {
			_, err := usersRepo.SelectUser(ctx, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: "",
				},
			})

			Expect(err).NotTo(BeNil())

			_, err = usersRepo.SelectUser(ctx, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: "123456789sdfgs",
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersRepo.SelectUser(c, &pbUsers.RequestSelectUser{
				Data: &pbUsers.RequestSelectUser_Data{
					UserId: "",
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Select Session User By Id", func() {
		It("SUCCESS", func() {
			sessionUser, err := usersRepo.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: dataProduct.UserId,
				},
			})

			Expect(err).To(BeNil())
			Expect(len(strings.Split(sessionUser.UserSession, "."))).To(Equal(3))
		})

		It("FAIL", func() {
			_, err := usersRepo.SelectSessionUserById(ctx, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: "what ever you want",
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := usersRepo.SelectSessionUserById(c, &pbUsers.RequestSelectSessionUserById{
				Data: &pbUsers.RequestSelectSessionUserById_Data{
					UserId: "what ever you want",
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
