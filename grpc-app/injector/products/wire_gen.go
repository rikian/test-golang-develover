// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package products

import (
	"go.uber.org/zap"
	"go/service1/grpc-app/protos/products"
	products4 "go/service1/grpc-app/service/products"
	products2 "go/service1/shared/repository/products"
	products3 "go/service1/shared/usecase/products"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeProductsService(db *gorm.DB, logger *zap.Logger) (products.ProductRPCServer, error) {
	productRepository := products2.NewProductsRepositoryImpl(db, logger)
	productUseCase := products3.NewProductUseCaseImpl(logger, productRepository)
	productRPCServer := products4.NewProductService(productUseCase, logger)
	return productRPCServer, nil
}
