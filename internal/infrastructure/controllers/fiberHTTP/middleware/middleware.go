package middleware

import (
	"golang-api-template/internal/domain/service"
	"golang-api-template/internal/infrastructure/controllers/safeobject"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Middleware interface {
	CreatePolicyFunc() func(*fiber.Ctx) error
	GetPolicy(ctx *fiber.Ctx) (*safeobject.Policy, error)
}

type middleware struct {
	auth service.Auth

	log *logrus.Entry
}

func NewMiddleware(
	auth service.Auth,
	log *logrus.Entry,
) Middleware {
	return &middleware{
		auth: auth,
		log:  log,
	}
}
