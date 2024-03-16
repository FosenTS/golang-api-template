package fiberHTTP

import (
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP/handlers"

	"github.com/gofiber/fiber/v2"
)

type FiberController interface {
	RegisterRoutes(r *fiber.App)
}

type fiberController struct {
	handlers *handlers.HandlerFiber
}

func NewFiberController(handlers *handlers.HandlerFiber) FiberController {
	return &fiberController{handlers: handlers}
}

func (fC *fiberController) RegisterRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")

	fC.handlers.RegisterGroup(apiGroup)
}
