package users

import (
	"context"
	pb "go/service1/grpc-app/protos/users"
	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
)

type UsersUseCase interface {
	SelectUser(context.Context, *pb.RequestSelectUser) (*table.User, error)
	SelectSessionUserById(context.Context, *pb.RequestSelectSessionUserById) (*entities.ResponSelectSessionUserById, error)
}
