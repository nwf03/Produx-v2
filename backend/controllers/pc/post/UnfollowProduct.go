package post

import (
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func UnfollowProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var User db.User
	db.DB.Preload("FollowedProducts").First(&User, "id = ?", userId)
	if User.ID == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	productName := strings.TrimSpace(strings.ReplaceAll(c.Params("product_name"), "%20", " "))

	var Product db.Product
	db.DB.First(&Product, "name = ?", productName)
	if Product.ID == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	err := db.DB.Model(&User).Association("FollowedProducts").Delete(&Product)
	if err != nil {
		return err
	}
	db.DB.Save(&User)

	return c.Status(200).JSON(fiber.Map{
		"message": "Product unfollowed",
	})
}



