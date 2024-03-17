package fiberHTTP

import (
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP/middleware"

	"github.com/gofiber/fiber/v2"
)

type FiberController interface {
	RegisterRoutes(r *fiber.App)
}

type fiberController struct {
	authHandler HandlerFiber
	apiHanlder  HandlerFiber

	middleware middleware.Middleware
}

type HandlerFiber interface {
	RegisterGroup(g fiber.Router)
}

func NewFiberController(authHandler HandlerFiber, middleware middleware.Middleware) FiberController {
	return &fiberController{authHandler: authHandler, middleware: middleware}
}

func (fC *fiberController) RegisterRoutes(app *fiber.App) {

	policyChecker := fC.middleware.CreatePolicyFunc()

	authGroup := app.Group("/auth")
	apiGroup := app.Group("/api", policyChecker)

	fC.authHandler.RegisterGroup(authGroup)
	fC.apiHanlder.RegisterGroup(apiGroup)
}
