package routes

import (
	"go-homework/controllers"

	"github.com/gofiber/fiber/v2"
)

func BookRoute(route fiber.Router) {
	route.Get("/books", controllers.IndexBook)
	route.Get("/books/:id", controllers.DetailBook)
	route.Post("/books", controllers.CreateBook)
	route.Put("/books/:id", controllers.UpdateBook)
	route.Delete("/books/:id", controllers.DeleteBook)
}
