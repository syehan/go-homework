package controllers

import (
	"encoding/base64"
	"go-homework/config"
	"go-homework/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {

	username := getUsername(c)
	var user models.User

	result := config.Database.Where("username = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

func getUsername(c *fiber.Ctx) string {
	username, _, _ := GetBasicAuth(c)
	return username
}

func GetBasicAuth(c *fiber.Ctx) (username, password string, ok bool) {
	auth := c.Get("Authorization")
	if auth == "" {
		return "", "", false
	}
	return parseBasicAuth(auth)
}

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !EqualFold(auth[:len(prefix)], prefix) {
		return "", "", false
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return "", "", false
	}
	cs := string(c)
	username, password, ok = strings.Cut(cs, ":")
	if !ok {
		return "", "", false
	}
	return username, password, true
}

func EqualFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if lower(s[i]) != lower(t[i]) {
			return false
		}
	}
	return true
}

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
