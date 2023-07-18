package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	AuthorID    int
	Author      Author
}
