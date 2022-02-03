package delete

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"tutorial/db"
)

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	var product db.Product
	fmt.Println("user id: ", userID, productID)
	err := db.DB.Where("id = ? and user_id = ?", productID, userID).First(&product).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Product not found"})
	}
	if err := product.Delete(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error deleting product"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Product deleted"})
}
