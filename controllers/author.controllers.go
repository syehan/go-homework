package controllers

import (
	"go-homework/config"
	"go-homework/models"

	"github.com/gofiber/fiber/v2"
)

func IndexAuthor(c *fiber.Ctx) error {
	var authors []models.Author

	config.Database.Find(&authors)
	return c.Status(200).JSON(authors)

}

func DetailAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	result := config.Database.Find(&author, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&author)
}

func CreateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&author)
	return c.Status(201).JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)
	id := c.Params("id")

	if err := c.BodyParser(author); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&author)
	return c.Status(200).JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	result := config.Database.Delete(&author, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
