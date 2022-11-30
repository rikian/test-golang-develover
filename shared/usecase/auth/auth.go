package auth

import (
	"context"
	pb "go/service1/grpc-app/protos/auth"
	"go/service1/shared/models/entities"
)

type AuthUseCase interface {
	RegisterUser(context.Context, *pb.RequestRegister) (*entities.ResponRegisterUser, error)
	LoginUser(context.Context, *pb.RequestLogin) (*entities.ResponLoginUser, error)
}
