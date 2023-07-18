package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Security(app *fiber.App) {
	app.Use(
		helmet.New(),
		csrf.New(),
		logger.New(),
	)
}
