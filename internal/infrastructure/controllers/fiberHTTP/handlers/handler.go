package handlers

import "golang-api-template/internal/domain/service"

type HandlerFiber struct {
	AuthService service.Auth
}

func NewHandlerFiber() *HandlerFiber {
	return &HandlerFiber{}
}
