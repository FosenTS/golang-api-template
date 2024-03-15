package gormDB

import (
	"errors"
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/domain/storage/gormDB/scheme"
	"golang-api-template/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository storage.User

type userRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewUserRepository(db *gorm.DB, log *logrus.Entry) (UserRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.User{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	err = db.AutoMigrate(&scheme.RefreshToken{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return &userRepository{db: db, log: log}, nil
}

func (aR *userRepository) InsertUser(user *dto.UserCreate) error {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.Create(&scheme.User{Login: user.Login, Password: user.Password, Permission: user.Permission})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (aR *userRepository) Find(user *scheme.User) (*entity.User, error) {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.User{
		ID:         user.ID,
		Login:      user.Login,
		Password:   user.Password,
		Permission: user.Permission,
	}, nil
}

func (aR *userRepository) DeleteByID(id uint) error {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.Delete(&scheme.User{ID: id})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}
