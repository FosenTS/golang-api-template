package authhandler

import (
	"golang-api-template/internal/domain/service"
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP"

	"github.com/sirupsen/logrus"
)

type handlerAuth struct {
	AuthService service.Auth

	log *logrus.Entry
}

func NewHandlerAuth(authService *service.Auth, log *logrus.Entry) fiberHTTP.HandlerFiber {
	return &handlerAuth{
		AuthService: *authService,
		log:         log,
	}
}
