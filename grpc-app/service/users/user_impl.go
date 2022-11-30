package users

import (
	"context"
	"go/service1/common"
	protos "go/service1/grpc-app/protos/auth"
	pb "go/service1/grpc-app/protos/users"
	"go/service1/shared/usecase/users"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type UsersServiceImpl struct {
	pb.UnimplementedUserRPCServer
	usersUseCase users.UsersUseCase
	logger       *zap.Logger
}

func NewUserServiceImpl(UsersUseCase users.UsersUseCase, logger *zap.Logger) pb.UserRPCServer {
	return &UsersServiceImpl{
		usersUseCase: UsersUseCase,
		logger:       logger,
	}
}

func (u *UsersServiceImpl) SelectUser(c context.Context, i *pb.RequestSelectUser) (*pb.ResponseSelectUser, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service select user")
		return nil, c.Err()
	default:
		result, err := u.usersUseCase.SelectUser(c, i)

		if err != nil {
			return nil, err
		}

		if err != nil {
			u.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		user := common.MapUserToPb(result)

		return &pb.ResponseSelectUser{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},

			Data: user,
		}, nil
	}
}

func (a *UsersServiceImpl) SelectSessionUserById(c context.Context, i *pb.RequestSelectSessionUserById) (*pb.ResponseSelectSessionUserById, error) {
	select {
	case <-c.Done():
		log.Print("cancelled from service select user by id")
		return nil, c.Err()
	default:
		res, err := a.usersUseCase.SelectSessionUserById(c, i)

		if err != nil {
			err := status.Errorf(3, err.Error())
			return nil, err
		}

		return &pb.ResponseSelectSessionUserById{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},
			Data: &pb.ResponseSelectSessionUserById_Data{
				UserSession: res.UserSession,
				RememberMe:  res.RememberMe,
			},
		}, nil
	}
}
