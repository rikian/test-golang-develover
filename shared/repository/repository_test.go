package repository

import (
	// repository
	"go/service1/shared/repository/auth"
	"go/service1/shared/repository/products"
	"go/service1/shared/repository/users"

	// protobuff
	pbAuth "go/service1/grpc-app/protos/auth"
	pbProducts "go/service1/grpc-app/protos/products"
	pbUsers "go/service1/grpc-app/protos/users"

	// entities
	"go/service1/shared/models/entities/table"

	"context"
	"log"
	"strings"

	"go/service1/common"
	"go/service1/config"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	dbConn       *gorm.DB
	logger       *zap.Logger
	authRepo     auth.AuthRepository
	usersRepo    users.UserRepository
	productsRepo products.ProductRepository
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

func TestMain(m *testing.M) {
	config.LoadEnvFile()
	dbConn = config.ConnectDB()
	logger = common.BuildLogger()
	authRepo = auth.NewAuthRepositoryImpl(dbConn, logger)
	usersRepo = users.NewUsersRepositoryImpl(dbConn, logger)
	productsRepo = products.NewProductsRepositoryImpl(dbConn, logger)

	m.Run()

	dbConn.Where("user_email = ?", dataUser.UserEmail).Delete(&table.User{})
}

func TestRepository(t *testing.T) {
	// ==========TEST PACKAGE AUTH REPOSITORY==========
	// register user
	registerUser, err := authRepo.RegisterUser(context.Background(), &pbAuth.RequestRegister{
		Data: dataUser,
	})

	require.Nil(t, err)
	require.NotNil(t, registerUser)

	log.Print("SUCCESS TEST REGISTER USER")

	// login user
	loginUser, err := authRepo.LoginUser(context.Background(), &pbAuth.RequestLogin{
		Data: &pbAuth.RequestLogin_Data{
			UserEmail:      dataUser.UserEmail,
			UserPassword:   dataUser.UserPassword,
			UserRememberMe: true,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, loginUser)

	dataProduct.UserId = loginUser.UserId

	if dataProduct.UserId == "" {
		t.Fatal()
	}

	log.Print("SUCCESS TEST LOGIN USER")

	// ---------------TEST PACKAGE Users REPOSITORY---------------------
	// select user
	getUser, err := usersRepo.SelectUser(context.Background(), &pbUsers.RequestSelectUser{
		Data: &pbUsers.RequestSelectUser_Data{
			UserId: loginUser.UserId,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, getUser)

	log.Print("SUCCESS TEST SELECT USER")

	// select user session by id
	selectSessionUserById, err := usersRepo.SelectSessionUserById(context.Background(), &pbUsers.RequestSelectSessionUserById{
		Data: &pbUsers.RequestSelectSessionUserById_Data{
			UserId: loginUser.UserId,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, selectSessionUserById)
	require.Equal(t, 3, len(strings.Split(selectSessionUserById.UserSession, ".")))

	log.Print("SUCCESS TEST SELECT USER SESSION BY ID")

	// ---------------TEST PACKAGE PRODUCTS REPOSITORY---------------------
	// insert product
	insertProduct, err := productsRepo.InsertProduct(context.Background(), &pbProducts.Request{
		Data: dataProduct,
	})

	require.Nil(t, err)
	require.NotNil(t, insertProduct)

	log.Print("SUCCESS TEST INSERT PRODUCT")

	// get all product
	getAllProduct, err := productsRepo.GetAllProduct(context.Background(), &pbProducts.RequestGetAllProduct{
		Data: &pbProducts.RequestGetAllProduct_Data{
			Limit: 5,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, getAllProduct)

	for _, product := range getAllProduct {
		if product.UserId == dataProduct.UserId &&
			product.ProductName == dataProduct.ProductName &&
			product.ProductInfo == dataProduct.ProductInfo &&
			product.CategoryName == dataProduct.ProductCategory &&
			product.ProductPrice == dataProduct.ProductPrice &&
			product.ProductSell == dataProduct.ProductSell &&
			product.ProductStock == dataProduct.ProductStock {
			dataProduct.ProductId = product.ProductId
			break
		}
	}

	if dataProduct.ProductId == "" {
		t.Fatal("failed generate id product")
	}

	log.Print("SUCCESS TEST GET ALL PRODUCT")

	// get product by id
	getProductById, err := productsRepo.GetProductById(context.Background(), &pbProducts.Request{
		Data: &pbProducts.Product{
			UserId:      loginUser.UserId,
			ProductId:   dataProduct.ProductId,
			ProductName: dataProduct.ProductName,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, getProductById)

	log.Print("SUCCESS TEST GET PRODUCT BY ID")

	// update product
	newDataProduct := &pbProducts.Product{
		UserId:          loginUser.UserId,
		ProductId:       dataProduct.ProductId,
		ProductName:     "Kue Tidak Basi",
		ProductCategory: "electronic",
		ProductInfo:     "Kue Basi tapi enak",
		ProductPrice:    100000,
		ProductSell:     80000,
		ProductStock:    20,
		ProductImage:    "example.com/images/userid/productid/images1.jpg",
	}
	updateProduct, err := productsRepo.UpdateProduct(context.Background(), &pbProducts.Request{
		Data: newDataProduct,
	})

	require.Nil(t, err)
	require.NotNil(t, updateProduct)

	// chect update product
	getProductByIdTest, err := productsRepo.GetProductById(context.Background(), &pbProducts.Request{
		Data: &pbProducts.Product{
			UserId:      loginUser.UserId,
			ProductId:   dataProduct.ProductId,
			ProductName: newDataProduct.ProductName,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, getProductByIdTest)

	if updateProduct.ProductName == getProductByIdTest.ProductName {
		log.Print("SUCCESS TEST UPDATE PRODUCT")
	} else {
		log.Print("FAILED TEST UPDATE PRODUCT")
		t.Fatal()
	}

	// delete product
	deleteProduct, err := productsRepo.DeleteProduct(context.Background(), &pbProducts.Request{
		Data: &pbProducts.Product{
			UserId:    loginUser.UserId,
			ProductId: dataProduct.ProductId,
		},
	})

	require.Nil(t, err)
	require.NotNil(t, deleteProduct)

	// chect delete product
	getProductByIdTest2, err := productsRepo.GetProductById(context.Background(), &pbProducts.Request{
		Data: &pbProducts.Product{
			UserId:      loginUser.UserId,
			ProductId:   dataProduct.ProductId,
			ProductName: newDataProduct.ProductName,
		},
	})

	require.NotNil(t, err)
	require.Nil(t, getProductByIdTest2)

	log.Print("SUCCESS TEST DELETE PRODUCT")
	log.Print("---------------------------")
	log.Print("==========SUCCESS TEST PACKAGE REPOSITORY==========")
}
