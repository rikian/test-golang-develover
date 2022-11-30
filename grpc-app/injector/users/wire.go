//go:build wireinject
// +build wireinject

package users

import (
	pbUsers "go/service1/grpc-app/protos/users"

	serviceUser "go/service1/grpc-app/service/users"
	repoUser "go/service1/shared/repository/users"

	ucUser "go/service1/shared/usecase/users"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitializeUsersService(db *gorm.DB, logger *zap.Logger) (pbUsers.UserRPCServer, error) {
	wire.Build(
		repoUser.NewUsersRepositoryImpl,
		ucUser.NewUsersUseCaseImpl,
		serviceUser.NewUserServiceImpl,
	)

	return &serviceUser.UsersServiceImpl{}, nil
}
