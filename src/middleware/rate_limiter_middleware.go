package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"abarobotics-test/src/constant"
)

func RateLimiterMiddleware(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:        constant.DefaultMdwRateLimiter,
		Expiration: constant.DefaultMdwRateLimiterDuration,
	}))
}
