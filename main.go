package main

import (
	"log"
	"os"

	"go-homework/config"
	"go-homework/middlewares"
	"go-homework/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error when loading .env")
		}
	}

	port := os.Getenv("GO_PORT")
	app := fiber.New()

	middlewares.BasicAuth(app)
	middlewares.Security(app)

	config.ConnectSqlite()

	api := app.Group("/api/v1")
	routes.AuthorRoute(api)
	routes.BookRoute(api)
	routes.UserRoute(api)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
