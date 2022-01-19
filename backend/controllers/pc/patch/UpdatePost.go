package patch

import (
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UpdatePost(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	postId := c.Params("post_id")
	newTitle := c.FormValue("title")
	newDescription := c.FormValue("description")
	newVersion := c.FormValue("version")
	if newTitle == "" && newDescription == "" && newVersion == "" {
		return c.Status(400).JSON(fiber.Map{"error": "title or description is required"})
	}

	if !mw.ValidFields(field) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid field"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	var product db.Product
	DB.First(&product, "name = ?", name)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	switch field {
	case "Suggestions":
		var suggestion db.Suggestion
		DB.First(&suggestion, "id = ?", postId)
		if suggestion.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Suggestion not found"})
		}
		if suggestion.UserID != uint(id) {
			return c.Status(403).JSON(fiber.Map{"error": "Not authorized"})
		}
		if newTitle != "" {
			suggestion.Title = newTitle
		}
		if newDescription != "" {
			suggestion.Description = newDescription
		}
		DB.Save(&suggestion)
		return c.Status(200).JSON(suggestion)
	case "Bugs":
		var bug db.Bug
		DB.First(&bug, "id = ?", postId)
		if bug.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Bug not found"})
		}
		if bug.UserID != uint(id) {
			return c.Status(403).JSON(fiber.Map{"error": "Not authorized"})
		}
		if newTitle != "" {
			bug.Title = newTitle
		}
		if newDescription != "" {
			bug.Description = newDescription
		}
		DB.Save(&bug)
		return c.Status(200).JSON(bug)
	case "Changelogs":
		var changelog db.Changelog

		DB.First(&changelog, "id = ?", postId)
		if changelog.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Changelog not found"})
		}
		if changelog.UserID != uint(id) {
			return c.Status(403).JSON(fiber.Map{"error": "Not authorized"})
		}
		if newTitle != "" {
			changelog.Title = newTitle
		}
		if newDescription != "" {
			changelog.Description = newDescription
		}
		if newVersion != "" {
			changelog.Version = newVersion
		}
		DB.Save(&changelog)
		return c.Status(200).JSON(changelog)

	}
	return c.Status(404).JSON(fiber.Map{"error": "field not found"})

}
