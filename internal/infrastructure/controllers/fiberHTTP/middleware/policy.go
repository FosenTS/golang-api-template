package middleware

import (
	"encoding/json"
	"golang-api-template/pkg/advancedlog"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) CreatePolicyFunc() func(*fiber.Ctx) error {
	logF := advancedlog.FunctionLog(m.log)
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		policy, err := m.auth.Policy(ctx.Context(), token)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		policyData, err := json.Marshal(policy)
		if err != nil {
			logF.Errorln(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		ctx.Set("policy", string(policyData))

		return ctx.Next()
	}
}
