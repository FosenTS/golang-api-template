package fiberHTTP

import "github.com/gofiber/fiber/v3"

type FiberController interface {
	RegisterRoutes(r *fiber.App)
}

type fiberController struct {
}

func (fC *fiberController) RegisterRoutes() {
}
