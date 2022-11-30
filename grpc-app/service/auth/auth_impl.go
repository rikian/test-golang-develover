package auth

import (
	"context"
	pb "go/service1/grpc-app/protos/auth"
	"go/service1/shared/usecase/auth"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type authServiceImpl struct {
	pb.UnimplementedAuthRPCServer
	authUseCase auth.AuthUseCase
	logger      *zap.Logger
}

func NewAuthService(authUseCase auth.AuthUseCase, logger *zap.Logger) pb.AuthRPCServer {
	return &authServiceImpl{
		authUseCase: authUseCase,
		logger:      logger,
	}
}

func (a *authServiceImpl) RegisterUser(ctx context.Context, i *pb.RequestRegister) (*pb.ResponseRegister, error) {
	select {
	case <-ctx.Done():
		log.Print("canceled from service register user")
		return nil, ctx.Err()
	default:
		res, err := a.authUseCase.RegisterUser(ctx, i)

		if err != nil {
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		return &pb.ResponseRegister{
			Info: &pb.Info{
				Code:   0,
				Status: "ok",
			},
			Data: &pb.ResponseRegister_Data{
				UserName:  res.UserName,
				UserEmail: res.UserEmail,
			},
		}, nil
	}
}

func (a *authServiceImpl) LoginUser(c context.Context, i *pb.RequestLogin) (*pb.ResponseLogin, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service login user")
		return nil, c.Err()
	default:
		res, err := a.authUseCase.LoginUser(c, i)

		if err != nil {
			err := status.Errorf(3, err.Error())
			return nil, err
		}

		return &pb.ResponseLogin{
			Info: &pb.Info{
				Code:   0,
				Status: "ok",
			},
			Data: &pb.ResponseLogin_Data{
				UserId:  res.UserId,
				Session: res.Session,
			},
		}, nil
	}
}
