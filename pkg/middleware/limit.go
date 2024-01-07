package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = limiter.New(limiter.Config{
	Max: 100,
	LimitReached: func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusTooManyRequests).Render("pages/error", map[string]any{
			"code":    fiber.StatusTooManyRequests,
			"message": "Too Many Requests, Please Try Again Later",
		})
	},
})
