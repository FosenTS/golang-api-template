package handlers

import (
	"github.com/sirupsen/logrus"
	"golang-api-template/internal/domain/service"
)

type HandlerFiber struct {
	AuthService service.Auth

	log *logrus.Entry
}

func NewHandlerFiber() *HandlerFiber {
	return &HandlerFiber{}
}
