package product

import (
	"golang-api-template/internal/application/config"
	"golang-api-template/internal/domain/service"
	"golang-api-template/pkg/ajwt"
	"golang-api-template/pkg/passlib"

	"github.com/sirupsen/logrus"
)

type Services struct {
	Auth service.Auth
}

func NewServices(
	storage *Storage,
	authConfig config.AuthConfig,
	log *logrus.Entry,
) *Services {
	hashManager := passlib.NewHashManager(authConfig.Salt)
	jwtManager := ajwt.NewJWTManager(hashManager, authConfig.SecretJWTKey, authConfig.JwtLiveTime, authConfig.RefreshLiveTime)

	authService := service.NewAuth(storage.User, storage.RefreshTokens, hashManager, jwtManager, log.WithField("location", "auth-service"))
	return &Services{
		Auth: authService,
	}
}
