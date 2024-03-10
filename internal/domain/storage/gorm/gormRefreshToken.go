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

type RefreshTokenRepository storage.RefreshToken

type refreshTokenRepository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewRefreshTokenRepository(db *gorm.DB, log *logrus.Entry) (RefreshTokenRepository, error) {
	logF := advancedlog.FunctionLog(log)

	err := db.AutoMigrate(&scheme.RefreshToken{})
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}
	return &refreshTokenRepository{db: db, log: log}, nil
}

func (rR *refreshTokenRepository) InsertRefreshToken(tCreate *dto.RefreshTokenCreate) (*entity.RefreshToken, error) {
	logF := advancedlog.FunctionLog(rR.log)
	refreshTokenF := scheme.RefreshToken{
		Token:              tCreate.Token,
		Login:              tCreate.Login,
		ExpirationTimeUnix: tCreate.ExpirationTimeUnix,
		CreateTimeUnix:     tCreate.CreateTimeUnix,
	}
	result := rR.db.First(&refreshTokenF)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.RefreshToken{
		ID:                 refreshTokenF.ID,
		Token:              refreshTokenF.Token,
		Login:              refreshTokenF.Login,
		ExpirationTimeUnix: refreshTokenF.ExpirationTimeUnix,
		CreateTimeUnix:     refreshTokenF.CreateTimeUnix,
	}, nil
}

func (rR *refreshTokenRepository) Find(refreshToken *scheme.RefreshToken) (*entity.RefreshToken, error) {
	logF := advancedlog.FunctionLog(rR.log)

	result := rR.db.First(refreshToken)
	if result.Error != nil {
		logF.Errorln(result.Error)
		return nil, result.Error
	}

	return &entity.RefreshToken{
		ID:                 refreshToken.ID,
		Token:              refreshToken.Token,
		Login:              refreshToken.Login,
		ExpirationTimeUnix: refreshToken.ExpirationTimeUnix,
		CreateTimeUnix:     refreshToken.CreateTimeUnix,
	}, nil
}
