package product

import (
	"golang-api-template/internal/application/config"
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP"
	apihandler "golang-api-template/internal/infrastructure/controllers/fiberHTTP/handlers/apiHandler"
	authhandler "golang-api-template/internal/infrastructure/controllers/fiberHTTP/handlers/authHandler"
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP/middleware"
	"golang-api-template/internal/infrastructure/controllers/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	fiberController fiberHTTP.FiberController
	metrics         metrics.Metrics
}

func NewController(gateway *Gateway, metricsConfig config.MetricsConfig, log *logrus.Entry) *Controller {
	middlew := middleware.NewMiddleware(gateway.Services.Auth, log.WithField("location", "middleware"))
	return &Controller{
		fiberController: fiberHTTP.NewFiberController(authhandler.NewHandlerAuth(gateway.Services.Auth, middlew, log.WithField("location", "handler-auth")), apihandler.NewHandlerApi(log.WithField("location", "handler-api")), middlew),
		metrics:         metrics.NewMetrics(metricsConfig, log.WithField("location", "metrics-listener")),
	}
}

func (c *Controller) ConfigureFiber(r *fiber.App) error {
	c.fiberController.RegisterRoutes(r)

	return nil
}

func (c *Controller) ListenMetrics() error {
	return c.metrics.Listen()
}
