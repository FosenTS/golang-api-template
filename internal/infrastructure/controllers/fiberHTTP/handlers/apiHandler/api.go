package apihandler

import "github.com/gofiber/fiber/v2"

func (h *handlerApi) RegisterGroup(g fiber.Router) {
	g.Post("/test", h.Test)
}

func (h *handlerApi) Test(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("test")
}
