//go:build wireinject
// +build wireinject

package auth

import (
	pbAuth "go/service1/grpc-app/protos/auth"
	serviceAuth "go/service1/grpc-app/service/auth"
	repoAuth "go/service1/shared/repository/auth"
	ucAuth "go/service1/shared/usecase/auth"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitializeAuthService(db *gorm.DB, logger *zap.Logger) (pbAuth.AuthRPCServer, error) {
	wire.Build(
		repoAuth.NewAuthRepositoryImpl,
		ucAuth.NewAuthUseCaseImpl,
		serviceAuth.NewAuthService,
	)

	return &serviceAuth.AuthService{}, nil
}
