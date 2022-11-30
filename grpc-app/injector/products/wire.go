//go:build wireinject
// +build wireinject

package products

import (
	pbProduct "go/service1/grpc-app/protos/products"
	serviceProduct "go/service1/grpc-app/service/products"
	repoProduct "go/service1/shared/repository/products"
	ucProduct "go/service1/shared/usecase/products"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitializeProductsService(db *gorm.DB, logger *zap.Logger) (pbProduct.ProductRPCServer, error) {
	wire.Build(
		repoProduct.NewProductsRepositoryImpl,
		ucProduct.NewProductUseCaseImpl,
		serviceProduct.NewProductService,
	)

	return &serviceProduct.ProductsServiceImpl{}, nil
}
