package controllers

import (
	"go-homework/config"
	"go-homework/models"

	"github.com/gofiber/fiber/v2"
)

func IndexBook(c *fiber.Ctx) error {
	var books []models.Book

	config.Database.Preload("Author").Find(&books)
	return c.Status(200).JSON(books)

}

func DetailBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := config.Database.Preload("Author").Find(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&book)
	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	id := c.Params("id")

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&book)
	return c.Status(200).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := config.Database.Delete(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
