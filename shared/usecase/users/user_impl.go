package users

import (
	"context"
	pb "go/service1/grpc-app/protos/users"
	"log"

	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
	"go/service1/shared/repository/users"

	"go.uber.org/zap"
)

type usersUseCaseImpl struct {
	UserRepo users.UserRepository
	logger   *zap.Logger
}

func NewUsersUseCaseImpl(userRepo users.UserRepository, logger *zap.Logger) UsersUseCase {
	return &usersUseCaseImpl{
		UserRepo: userRepo,
		logger:   logger,
	}
}

func (u *usersUseCaseImpl) SelectUser(c context.Context, i *pb.RequestSelectUser) (*table.User, error) {
	select {
	case <-c.Done():
		log.Print("cancelled from usecase select user")
		return nil, c.Err()
	default:
		res, err := u.UserRepo.SelectUser(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func (u *usersUseCaseImpl) SelectSessionUserById(c context.Context, i *pb.RequestSelectSessionUserById) (*entities.ResponSelectSessionUserById, error) {
	select {
	case <-c.Done():
		log.Print("cancelled from usecase select session user by id")
		return nil, c.Err()
	default:
		res, err := u.UserRepo.SelectSessionUserById(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
