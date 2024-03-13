package handlers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/pkg/advancedlog"
)

func (h *HandlerFiber) RegisterGroup(g *fiber.Group) {
	g.Post("/login", h.Login)
	g.Post("/register", h.Register)
}

func (h *HandlerFiber) Login(c fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	body := c.Body()
	var login *dto.Login
	err := json.Unmarshal(body, login)
	if err != nil {
		logF.Warnln(err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid login format")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	user, tokens, err := h.AuthService.Login(ctx, login)
	if err != nil {
		logF.Errorln(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return
}

func (h *HandlerFiber) Register(c fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	body := c.Body()
	var register *dto.UserCreate
	err := json.Unmarshal(body, register)
	if err != nil {
		logF.Warnln(err)
		return c.Status(fiber.StatusBadRequest).SendString("invalid register format")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = h.AuthService.Register(ctx, register)
	if err != nil {
		logF.Errorln(err)
		return c.Status(fiber.StatusInternalServerError).SendString("user not created")
	}

	return c.SendStatus(fiber.StatusOK)
}
