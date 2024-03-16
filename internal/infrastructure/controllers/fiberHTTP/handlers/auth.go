package handlers

import (
	"context"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/pkg/advancedlog"

	"github.com/gofiber/fiber/v2"
)

func (h *HandlerFiber) RegisterGroup(g fiber.Router) {
	g.Post("/login", h.Login)
	g.Post("/register", h.Register)
}

func (h *HandlerFiber) Login(c *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	login := new(dto.Login)
	if err := c.BodyParser(login); err != nil {
		logF.Errorln(err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid login format")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tokens, err := h.AuthService.Login(ctx, login)
	if err != nil {
		logF.Errorln(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Status(fiber.StatusOK).JSON(tokens)
}

func (h *HandlerFiber) Register(c *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	register := new(dto.UserCreate)
	if err := c.BodyParser(register); err != nil {
		logF.Errorln(err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid register format")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := h.AuthService.Register(ctx, register)
	if err != nil {
		logF.Errorln(err)
		return c.Status(fiber.StatusInternalServerError).SendString("user not created")
	}

	return c.SendStatus(fiber.StatusOK)
}
