package handlers

import (
	"golang-api-template/internal/domain/service"

	"github.com/sirupsen/logrus"
)

type HandlerFiber struct {
	AuthService service.Auth

	log *logrus.Entry
}

func NewHandlerFiber() *HandlerFiber {
	return &HandlerFiber{}
}
