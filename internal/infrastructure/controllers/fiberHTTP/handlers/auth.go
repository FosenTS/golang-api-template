package handlers

import "github.com/gofiber/fiber/v3"

func (h *HandlerFiber) RegisterGroup(g *fiber.Group) {
	g.Post("/login", h.Login)
	g.Post("/register", h.Register)
}

func (h *HandlerFiber) Login(ctx fiber.Ctx) error {

}

func (h *HandlerFiber) Register(ctx fiber.Ctx) error {

}
