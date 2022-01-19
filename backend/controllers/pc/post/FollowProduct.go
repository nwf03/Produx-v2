package post

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func FollowProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var User db.User
	DB.Preload("FollowedProducts").First(&User, "name = ?", name)
	if User.ID == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	productName := c.FormValue("product_name")

	var Product db.Product
	DB.First(&Product, "name = ?", productName)
	if Product.ID == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	DB.Model(&User).Association("FollowedProducts").Append(&Product)
	DB.Save(&User)
	return c.Status(200).JSON(fiber.Map{
		"message": "Product followed",
	})
}