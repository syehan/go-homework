package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}
