package users

import (
	"context"
	pb "go/service1/grpc-app/protos/users"
)

type UsersService interface {
	SelectUser(context.Context, *pb.RequestSelectUser) (*pb.ResponseSelectUser, error)
	SelectSessionUserById(context.Context, *pb.RequestSelectSessionUserById) (*pb.ResponseSelectSessionUserById, error)
}
