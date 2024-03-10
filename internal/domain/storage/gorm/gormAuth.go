package gorm

import (
	"github.com/sirupsen/logrus"
	"golang-api-template/internal/domain/entity"
	"golang-api-template/internal/domain/storage"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/internal/domain/storage/gorm/scheme"
	"golang-api-template/pkg/advancedlog"
	"gorm.io/gorm"
)

type AuthRepository storage.Auth

type authRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewAuthRepository(db *gorm.DB, log *logrus.Entry) (AuthRepository, error) {
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
	return &authRepository{db: db, log: log}, nil
}

func (aR *authRepository) InsertUser(user *dto.UserCreate) error {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.Create(&scheme.User{Login: user.Login, Password: user.Password, Permission: user.Permission})
	if result.Error != nil {
		logF.Errorln(result.Error)
		return result.Error
	}

	return nil
}

func (aR *authRepository) Find(user *scheme.User) (*entity.User, error) {
	logF := advancedlog.FunctionLog(aR.log)
	result := aR.db.First(user)
	if result.Error != nil {
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
