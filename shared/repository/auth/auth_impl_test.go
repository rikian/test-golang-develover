package auth

import (
	"context"
	"go/service1/common"
	"go/service1/config"
	pbAuth "go/service1/grpc-app/protos/auth"
	"go/service1/shared/models/entities/table"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	authRepo AuthRepository
	dbConn   *gorm.DB
	logger   *zap.Logger
	ctx      context.Context
)

var dataUser = &pbAuth.RequestRegister_Data{
	UserEmail:    "rikianfaisal@gmail.com",
	UserName:     "Rikian Faisal",
	UserPassword: "123456",
}

func TestAuthRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Products Repo Suite")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	dbConn = config.ConnectDB()
	logger = common.BuildLogger()
	authRepo = NewAuthRepositoryImpl(dbConn, logger)
})

var _ = AfterSuite(func() {
	// delete user register
	dbConn.Where("user_email = ?", dataUser.UserEmail).Delete(&table.User{})
})

var _ = Describe("Checking Products Repository", Ordered, func() {
	Context("Register And Login User", func() {
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
		})

		It("FAIL", func() {
			_, err := authRepo.RegisterUser(ctx, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).NotTo(BeNil())

			// check if user was stored in db or not
			_, err = authRepo.LoginUser(ctx, &pbAuth.RequestLogin{
				Data: &pbAuth.RequestLogin_Data{
					UserEmail:      "",
					UserPassword:   dataUser.UserPassword,
					UserRememberMe: false,
				},
			})

			Expect(err).NotTo(BeNil())

			_, err = authRepo.LoginUser(ctx, &pbAuth.RequestLogin{
				Data: &pbAuth.RequestLogin_Data{
					UserEmail:      dataUser.UserEmail,
					UserPassword:   "",
					UserRememberMe: false,
				},
			})

			Expect(err).NotTo(BeNil())

		})

		It("CANCEL REGISTER", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authRepo.RegisterUser(c, &pbAuth.RequestRegister{
				Data: dataUser,
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})

		It("CANCEL LOGIN", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := authRepo.LoginUser(c, &pbAuth.RequestLogin{
				Data: &pbAuth.RequestLogin_Data{
					UserEmail:      dataUser.UserEmail,
					UserPassword:   dataUser.UserPassword,
					UserRememberMe: false,
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
