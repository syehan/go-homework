package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"syehan": "ganteng",
			"admin":  "admin123",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "syehan" && pass == "ganteng" {
				return true
			}
			if user == "admin" && pass == "admin123" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			exception := map[string]string{"message": "Gaboleh masuk bro"}
			return c.Status(401).JSON(exception)
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))
}
