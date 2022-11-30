package auth

import (
	"context"
	"errors"
	pb "go/service1/grpc-app/protos/auth"
	"go/service1/shared/models/entities"
	"go/service1/shared/repository/auth"
	"log"

	"go.uber.org/zap"
)

type authUseCaseImpl struct {
	AuthRepo auth.AuthRepository
	logger   *zap.Logger
}

func NewAuthUseCaseImpl(authRepo auth.AuthRepository, logger *zap.Logger) AuthUseCase {
	return &authUseCaseImpl{
		AuthRepo: authRepo,
		logger:   logger,
	}
}

func (a *authUseCaseImpl) RegisterUser(c context.Context, i *pb.RequestRegister) (*entities.ResponRegisterUser, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase register user")
		return nil, c.Err()
	default:
		res, err := a.AuthRepo.RegisterUser(c, i)

		if err != nil {
			a.logger.Info(err.Error())
			return nil, errors.New(err.Error())
		}

		return res, nil
	}
}

func (a *authUseCaseImpl) LoginUser(c context.Context, i *pb.RequestLogin) (*entities.ResponLoginUser, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase login user")
		return nil, c.Err()
	default:
		res, err := a.AuthRepo.LoginUser(c, i)

		if err != nil {
			a.logger.Info(err.Error())
			return nil, errors.New(err.Error())
		}

		return res, nil
	}
}
