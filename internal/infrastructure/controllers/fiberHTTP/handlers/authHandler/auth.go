package authhandler

import (
	"golang-api-template/internal/domain/storage/dto"
	"golang-api-template/pkg/advancedlog"

	"github.com/gofiber/fiber/v2"
)

func (h *handlerAuth) RegisterGroup(g fiber.Router) {
	g.Post("/login", h.Login)
	g.Post("/register", h.Register)
	g.Get("/check", h.Check)
	g.Post("/refresh", h.Refresh)
}

func (h *handlerAuth) Login(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	login := new(dto.Login)
	if err := ctx.BodyParser(login); err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid login format")
	}

	tokens, err := h.AuthService.Login(ctx.Context(), login)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Status(fiber.StatusOK).JSON(tokens)
}

func (h *handlerAuth) Refresh(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	token := ctx.Query("refresh")

	pairTokens, err := h.AuthService.Refresh(ctx.Context(), token)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Status(fiber.StatusOK).JSON(pairTokens)
}

func (h *handlerAuth) Check(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)
	token := ctx.Get("Authorization")

	user, err := h.AuthService.Check(ctx.Context(), token)
	if err != nil {
		logF.Errorln(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *handlerAuth) Register(ctx *fiber.Ctx) error {
	logF := advancedlog.FunctionLog(h.log)

	register := new(dto.UserCreate)
	if err := ctx.BodyParser(register); err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid register format")
	}

	err := h.AuthService.Register(ctx.Context(), register)
	if err != nil {
		logF.Errorln(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("user not created")
	}

	return ctx.SendStatus(fiber.StatusOK)
}
