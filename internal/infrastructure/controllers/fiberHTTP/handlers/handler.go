package handlers

import (
	"golang-api-template/internal/domain/service"

	"github.com/sirupsen/logrus"
)

type HandlerFiber struct {
	AuthService service.Auth

	log *logrus.Entry
}

func NewHandlerFiber(authService *service.Auth, log *logrus.Entry) *HandlerFiber {
	return &HandlerFiber{
		AuthService: *authService,
		log:         log.WithField("location", "handler-fiber"),
	}
}
