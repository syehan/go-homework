package routes

import (
	"go-homework/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	route.Get("/me", controllers.Me)
}
