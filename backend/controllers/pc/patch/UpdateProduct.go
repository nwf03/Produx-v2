package patch

import (
	"fmt"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//var DB = controllers.DB

func UpdateProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	productName := c.Params("name")
	var product db.Product
	fmt.Println("product name: ", productName)
	db.DB.First(&product, "user_id = ? and name = ?", id, strings.ReplaceAll(productName, "%20", " "))
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
	}
	newName := c.FormValue("name")
	newDescription := c.FormValue("description")
	if newName == "" && newDescription == "" {
		return c.Status(400).JSON(fiber.Map{"message": "name or description is required"})
	}
	if newName != "" {
		product.Name = newName
	}
	if newDescription != "" {
		product.Description = newDescription
	}
	db.DB.Save(&product)
	return c.Status(200).JSON(product)
}
