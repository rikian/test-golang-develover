package products

import (
	"context"
	"go/service1/common"
	"go/service1/config"
	pbAuth "go/service1/grpc-app/protos/auth"
	pbProducts "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
	"go/service1/shared/repository/auth"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	productsRepo ProductRepository
	authRepo     auth.AuthRepository
	dbConn       *gorm.DB
	logger       *zap.Logger
	ctx          context.Context
)

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

var newDataProduct = &pbProducts.Product{
	ProductName:     "Kue tidak basi",
	ProductCategory: "electronic", // do not edit (constrains)
	ProductInfo:     "Kue Basi tapi enak",
	ProductPrice:    99999,
	ProductSell:     99999,
	ProductStock:    999,
	ProductImage:    "example.com/images/userid/productid/images1.jpg",
}

func TestProductsRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Products Repo Suite")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	dbConn = config.ConnectDB()
	logger = common.BuildLogger()
	productsRepo = NewProductsRepositoryImpl(dbConn, logger)
	authRepo = auth.NewAuthRepositoryImpl(dbConn, logger)

	// register user
	_, err := authRepo.RegisterUser(ctx, &pbAuth.RequestRegister{
		Data: dataUser,
	})

	Expect(err).To(BeNil())

	// login for get user id
	resLogin, err := authRepo.LoginUser(ctx, &pbAuth.RequestLogin{
		Data: &pbAuth.RequestLogin_Data{
			UserEmail:      dataUser.UserEmail,
			UserPassword:   dataUser.UserPassword,
			UserRememberMe: false,
		},
	})

	Expect(err).To(BeNil())
	Expect(resLogin.UserId).NotTo(Equal(""))

	// passing user id to product data
	dataProduct.UserId = resLogin.UserId
})

var _ = AfterSuite(func() {
	// delete user register
	dbConn.Where("user_email = ?", dataUser.UserEmail).Delete(&table.User{})
})

var _ = Describe("Checking Products Repository", Ordered, func() {
	Context("Insert Product", func() {
		It("SUCCESS", func() {
			insertProduct, err := productsRepo.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).To(BeNil())

			getProduct, err := productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   insertProduct.ProductId,
					ProductName: insertProduct.ProductName,
				},
			})

			Expect(err).To(BeNil())
			Expect(getProduct.ProductName).To(Equal(dataProduct.ProductName))
			Expect(getProduct.CategoryName).To(Equal(dataProduct.ProductCategory))
			Expect(getProduct.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(getProduct.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(getProduct.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(getProduct.ProductStock).To(Equal(dataProduct.ProductStock))

			// passing product id to data product for query
			dataProduct.ProductId = insertProduct.ProductId
		})

		It("FAIL", func() {
			_, err := productsRepo.InsertProduct(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsRepo.InsertProduct(c, &pbProducts.Request{
				Data: &pbProducts.Product{},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Get All Product", func() {
		It("SUCCESS", func() {
			getAllProduct, err := productsRepo.GetAllProduct(ctx, &pbProducts.RequestGetAllProduct{
				Data: &pbProducts.RequestGetAllProduct_Data{
					Limit: 5,
				},
			})

			Expect(err).To(BeNil())

			validProduct := false

			for _, product := range getAllProduct {
				if product.ProductId == dataProduct.ProductId && product.UserId == dataProduct.UserId {
					Expect(product.ProductName).To(Equal(dataProduct.ProductName))
					Expect(product.ProductInfo).To(Equal(dataProduct.ProductInfo))
					Expect(product.CategoryName).To(Equal(dataProduct.ProductCategory))
					Expect(product.ProductPrice).To(Equal(dataProduct.ProductPrice))
					Expect(product.ProductSell).To(Equal(dataProduct.ProductSell))
					Expect(product.ProductStock).To(Equal(dataProduct.ProductStock))

					validProduct = true

					break
				}
			}

			Expect(validProduct).To(Equal(true))
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsRepo.GetAllProduct(c, &pbProducts.RequestGetAllProduct{
				Data: &pbProducts.RequestGetAllProduct_Data{
					Limit: 5,
				},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Get Product By Id", func() {
		It("SUCCESS", func() {
			getProductById, err := productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					UserId:      dataProduct.UserId,
					ProductId:   dataProduct.ProductId,
					ProductName: dataProduct.ProductName,
				},
			})

			Expect(err).To(BeNil())
			Expect(getProductById.ProductName).To(Equal(dataProduct.ProductName))
			Expect(getProductById.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(getProductById.CategoryName).To(Equal(dataProduct.ProductCategory))
			Expect(getProductById.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(getProductById.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(getProductById.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("NOT FOUND", func() {
			_, err := productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   "",
					ProductName: dataProduct.ProductName,
				},
			})

			Expect(err).NotTo(BeNil())

			_, err = productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   dataProduct.ProductId,
					ProductName: "",
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsRepo.GetProductById(c, &pbProducts.Request{
				Data: &pbProducts.Product{},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Update Product", func() {
		It("SUCCESS", func() {
			newDataProduct.ProductId = dataProduct.ProductId
			newDataProduct.UserId = dataProduct.UserId

			_, err := productsRepo.UpdateProduct(ctx, &pbProducts.Request{
				Data: newDataProduct,
			})

			Expect(err).To(BeNil())

			getProductByIdTest, err := productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   dataProduct.ProductId,
					ProductName: newDataProduct.ProductName,
				},
			})

			Expect(err).To(BeNil())
			Expect(getProductByIdTest.ProductName).To(Equal(newDataProduct.ProductName))
			Expect(getProductByIdTest.ProductInfo).To(Equal(newDataProduct.ProductInfo))
			Expect(getProductByIdTest.CategoryName).To(Equal(newDataProduct.ProductCategory))
			Expect(getProductByIdTest.ProductPrice).To(Equal(newDataProduct.ProductPrice))
			Expect(getProductByIdTest.ProductSell).To(Equal(newDataProduct.ProductSell))
			Expect(getProductByIdTest.ProductStock).To(Equal(newDataProduct.ProductStock))
		})

		It("FAIL", func() {
			// get data product by id using old data
			_, err := productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   dataProduct.ProductId,
					ProductName: dataProduct.ProductName,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsRepo.UpdateProduct(c, &pbProducts.Request{
				Data: &pbProducts.Product{},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Delete Product", func() {
		It("SUCCESS", func() {
			_, err := productsRepo.DeleteProduct(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					UserId:    dataProduct.UserId,
					ProductId: dataProduct.ProductId,
				},
			})

			// chect with old data
			Expect(err).To(BeNil())
			_, err = productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   dataProduct.ProductId,
					ProductName: dataProduct.ProductName,
				},
			})

			Expect(err).NotTo(BeNil())

			// chect with new data
			_, err = productsRepo.GetProductById(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					ProductId:   dataProduct.ProductId,
					ProductName: newDataProduct.ProductName,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("FAIL", func() {
			_, err := productsRepo.DeleteProduct(ctx, &pbProducts.Request{
				Data: &pbProducts.Product{
					UserId:    dataProduct.UserId,
					ProductId: dataProduct.ProductId,
				},
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsRepo.DeleteProduct(c, &pbProducts.Request{
				Data: &pbProducts.Product{},
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
