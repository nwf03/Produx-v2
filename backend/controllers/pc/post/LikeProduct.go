package post

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LikeProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var User db.User
	db.DB.Preload("LikedProducts").Where("name = ?", name).First(&User)
	var Product db.Product
	db.DB.Preload("UserLikes").Where("id = ?", c.FormValue("id")).First(&Product)
	if Product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	for _, user := range Product.UserLikes {
		if user.Name == name {
			return c.Status(400).JSON(fiber.Map{
				"message": "You already liked this product",
			})
		}
	}
	err := db.DB.Model(&User).Association("LikedProducts").Append(&Product)
	if err != nil {
		return err
	}

	Product.LikesCount = Product.LikesCount + 1

	db.DB.Save(&Product)
	return c.Status(200).JSON(fiber.Map{
		"message": "Product liked",
	})

}
