package patch

import (
	"tutorial/controllers"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var DB = controllers.DB


func UpdateProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	var product db.Product
	DB.First(&product, "user_id", id)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	newName := c.FormValue("name")
	newDescription := c.FormValue("description")
	if newName == "" && newDescription == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name or description is required"})
	}
	if newName != "" {
		product.Name = newName
	}
	if newDescription != "" {
		product.Description = newDescription
	}
	DB.Save(&product)
	return c.Status(200).JSON(product)
}
