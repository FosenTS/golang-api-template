package product

import "github.com/gofiber/fiber/v3"

type Endpoint struct {
	*Controller
}

func NewEndpoint(controller *Controller) *Endpoint {
	return &Endpoint{Controller: controller}
}

func (e *Endpoint) ConfigureFiber(r *fiber.App) {
	e.ConfigureFiber(r)
}

func (e *Endpoint) ListenMetrics() error {
	err := e.ListenMetrics()
	if err != nil {
		return err
	}

	return nil
}
