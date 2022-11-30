package users

import (
	"context"
	"errors"
	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
	"log"

	pb "go/service1/grpc-app/protos/users"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUsersRepositoryImpl(db *gorm.DB, logger *zap.Logger) UserRepository {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (u *UserRepositoryImpl) SelectUser(c context.Context, i *pb.RequestSelectUser) (*table.User, error) {
	select {
	case <-c.Done():
		log.Print("canceled from user repo select user")
		return nil, c.Err()
	default:
		var user *table.User = &table.User{}

		err := u.db.Model(user).
			Preload("UserStatus").
			Preload("Products").
			Where("user_id = ?", i.Data.UserId).
			Find(user).Error

		if err != nil {
			log.Print(err.Error())
			return nil, err
		}

		if user.UserId == "" {
			return nil, errors.New("data not found")
		}

		return user, nil
	}
}

func (u *UserRepositoryImpl) SelectSessionUserById(c context.Context, i *pb.RequestSelectSessionUserById) (*entities.ResponSelectSessionUserById, error) {
	select {
	case <-c.Done():
		log.Print("canceled from user repo select user")
		return nil, c.Err()
	default:
		user := &table.User{}
		sessionUser := u.db.Select("user_session", "remember_me").
			Where("user_id = ?", i.Data.UserId).
			First(user)

		if sessionUser.Error != nil {
			u.logger.Info(sessionUser.Error.Error())
			return nil, sessionUser.Error
		}

		if sessionUser.RowsAffected != 1 {
			u.logger.Info("return row affected != 1")
			return nil, errors.New("return row affected != 1")
		}

		return &entities.ResponSelectSessionUserById{
			UserSession: user.UserSession,
			RememberMe:  user.RememberMe,
		}, nil
	}
}
