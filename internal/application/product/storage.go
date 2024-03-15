package product

import (
	"golang-api-template/internal/domain/storage"
	"golang-api-template/internal/domain/storage/gormDB"
	"golang-api-template/pkg/advancedlog"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Storage struct {
	User          storage.User
	RefreshTokens storage.RefreshToken
}

func NewStorage(db *gorm.DB, log *logrus.Entry) (*Storage, error) {
	logF := advancedlog.FunctionLog(log)
	userStorage, err := gormDB.NewUserRepository(db, log.WithField("location", "gorm-user-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	refreshTokenStorage, err := gormDB.NewRefreshTokenRepository(db, log.WithField("location", "gorm-refresh-token-repository"))
	if err != nil {
		logF.Errorln(err)
		return nil, err
	}

	return &Storage{
		User:          userStorage,
		RefreshTokens: refreshTokenStorage,
	}, nil
}
