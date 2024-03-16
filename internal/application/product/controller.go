package product

import (
	"golang-api-template/internal/application/config"
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP"
	"golang-api-template/internal/infrastructure/controllers/fiberHTTP/handlers"
	"golang-api-template/internal/infrastructure/controllers/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	fiberController fiberHTTP.FiberController
	metrics         metrics.Metrics
}

func NewController(gateway *Gateway, metricsConfig config.MetricsConfig, log *logrus.Entry) *Controller {
	return &Controller{
		fiberController: fiberHTTP.NewFiberController(handlers.NewHandlerFiber(&gateway.Services.Auth, log.WithField("location", "handler-fiber"))),
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
