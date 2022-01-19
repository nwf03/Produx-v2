package delete

import (
	"strings"
	"tutorial/controllers"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var DB = controllers.DB

func DeletePost(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	postId := c.Params("post_id")
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
		DB.Unscoped().Delete(&suggestion)
		return c.Status(200).JSON(fiber.Map{"message": "Suggestion deleted"})
	case "Bugs":
		var bug db.Bug
		DB.First(&bug, "id = ?", postId)
		if bug.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Bug not found"})
		}
		if bug.UserID != uint(id) {
			return c.Status(403).JSON(fiber.Map{"error": "Not authorized"})
		}
		DB.Unscoped().Delete(&bug)
		return c.Status(200).JSON(fiber.Map{"message": "Bug deleted"})
	case "Changelogs":
		var changelog db.Changelog

		DB.First(&changelog, "id = ?", postId)
		if changelog.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Changelog not found"})
		}
		if changelog.UserID != uint(id) {
			return c.Status(403).JSON(fiber.Map{"error": "Not authorized"})
		}
		DB.Unscoped().Delete(&changelog)
		return c.Status(200).JSON(fiber.Map{"message": "Changelog deleted"})
	}
	return c.Status(404).JSON(fiber.Map{"error": "field not found"})
}
