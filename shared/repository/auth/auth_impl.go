package auth

import (
	"context"
	"errors"
	h "go/service1/shared/http"
	"go/service1/shared/models/entities"
	"go/service1/shared/models/entities/table"
	"go/service1/shared/utils"
	"os"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	pb "go/service1/grpc-app/protos/auth"
)

type AuthRepoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAuthRepositoryImpl(db *gorm.DB, logger *zap.Logger) AuthRepository {
	return &AuthRepoImpl{
		db:     db,
		logger: logger,
	}
}

func (a *AuthRepoImpl) RegisterUser(c context.Context, q *pb.RequestRegister) (*entities.ResponRegisterUser, error) {
	// log.Print("processing repository register...")
	// init transaction
	tx := a.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			a.logger.Info("rollback error")
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		a.logger.Info(tx.Error.Error())
		return nil, tx.Error
	}

	select {
	case <-c.Done():
		// log.Println("canceled from repository register user")
		tx.Rollback()
		return nil, c.Err()
	default:
		var time string = time.Now().Format("20060102150405")
		user := &table.User{
			UserId:       uuid.New().String(),
			UserEmail:    q.Data.UserEmail,
			UserName:     q.Data.UserName,
			UserPassword: q.Data.UserPassword,
			UserImage:    "default.svg",
			UserSession:  "12345",
			UserStatusId: 1,
			CreatedDate:  time,
			LastUpdate:   time,
		}

		// logic
		createUser := tx.Create(&user)

		if createUser.Error != nil {
			a.logger.Info(createUser.Error.Error())
			tx.Rollback()
			return nil, createUser.Error
		}

		if createUser.RowsAffected == 0 {
			a.logger.Info("return row affected == 0")
			tx.Rollback()
			return nil, errors.New("return row affected == 0")
		}

		// create directory image
		debug := os.Getenv("DEBUG")

		if debug == "false" {
			httpRequest := h.NewHttpRequest()
			createDirMedia, err := httpRequest.CreateDirectoryImage(user.UserId)

			if err != nil {
				a.logger.Info(err.Error())
				tx.Rollback()
				return nil, err
			}

			if createDirMedia.Message != "ok" {
				a.logger.Info(createDirMedia.Message)
				tx.Rollback()
				return nil, errors.New(createDirMedia.Message)
			}
		}

		// comit
		comit := tx.Commit()
		if comit.Error != nil {
			a.logger.Info(comit.Error.Error())
			return nil, comit.Error
		}

		return &entities.ResponRegisterUser{
			UserName:  q.Data.UserName,
			UserEmail: q.Data.UserEmail,
		}, nil
	}
}

func (a *AuthRepoImpl) LoginUser(c context.Context, q *pb.RequestLogin) (*entities.ResponLoginUser, error) {
	var user *table.User = &table.User{}
	// begin the transaction
	tx := a.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			a.logger.Info("transaction failed")
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		a.logger.Info(tx.Error.Error())
		tx.Rollback()
		return nil, tx.Error
	}

	select {
	case <-c.Done():
		a.logger.Info("auth-repo-login : request canceled")
		tx.Rollback()
		return nil, c.Err()
	default:
		// query
		selectUser := tx.Select("user_id").
			Where("user_email = ? AND user_password = ?", q.Data.UserEmail, q.Data.UserPassword).
			First(user)

		if selectUser.Error != nil {
			a.logger.Info(selectUser.Error.Error())
			tx.Rollback()
			return nil, selectUser.Error
		}

		if selectUser.RowsAffected != 1 {
			a.logger.Info("rows affected not equal 1")
			tx.Rollback()
			return nil, errors.New("rows affected not equal 1")
		}

		var timeDuration int

		if q.Data.UserRememberMe {
			timeDuration = 31536000
		} else {
			timeDuration = 1800
		}

		sessionUser, err := utils.EncryptSession(user.UserId, timeDuration)

		if err != nil {
			a.logger.Info(err.Error())
			tx.Rollback()
			return nil, err
		}

		user.UserSession = sessionUser
		user.RememberMe = q.Data.UserRememberMe
		user.LastUpdate = time.Now().Format("20060102150405")

		insertSession := a.db.Model(user).
			Select("UserSession", "LastUpdate", "RememberMe").Updates(user).
			Where("user_email = ? AND user_password = ?", q.Data.UserEmail, q.Data.UserPassword)

		if insertSession.Error != nil {
			a.logger.Info(err.Error())
			tx.Rollback()
			return nil, err
		}

		if insertSession.RowsAffected != 1 {
			a.logger.Info("rows affected not equal 1")
			tx.Rollback()
			return nil, errors.New("rows affected not equal 1")
		}

		// comit
		comit := tx.Commit()

		if comit.Error != nil {
			a.logger.Info(comit.Error.Error())
			return nil, comit.Error
		}

		return &entities.ResponLoginUser{
			UserId:  user.UserId,
			Session: user.UserSession,
		}, nil
	}

}
