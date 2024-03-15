package fiberHTTP

import (
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP/handlers"

	"github.com/gofiber/fiber/v3"
)

type FiberController interface {
	RegisterRoutes(r *fiber.App)
}

type fiberController struct {
	handlers *handlers.HandlerFiber
}

func (fC *fiberController) RegisterRoutes(g *fiber.Group) {
	fC.handlers.RegisterGroup(g)
}
