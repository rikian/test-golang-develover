package auth

import (
	"context"
	pb "go/service1/grpc-app/protos/auth"
)

type AuthService interface {
	RegisterUser(context.Context, *pb.RequestRegister) (*pb.ResponseRegister, error)
	LoginUser(context.Context, *pb.RequestLogin) (*pb.ResponseLogin, error)
}
