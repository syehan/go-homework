package routes

import (
	"go-homework/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthorRoute(route fiber.Router) {
	route.Get("/authors", controllers.IndexAuthor)
	route.Get("/authors/:id", controllers.DetailAuthor)
	route.Post("/authors", controllers.CreateAuthor)
	route.Put("/authors/:id", controllers.UpdateAuthor)
	route.Delete("/authors/:id", controllers.DeleteAuthor)
}
