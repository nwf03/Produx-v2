package post

import (
	"fmt"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LikeProduct(c *fiber.Ctx) error{
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var User db.User
	DB.Preload("LikedProducts").Where("name = ?", name).First(&User)
	var Product db.Product
	DB.Preload("UserLikes").Where("id = ?", c.FormValue("id")).First(&Product)
	fmt.Println("Product name: ", Product.Name)
	if Product.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	for _, user := range Product.UserLikes{
		if user.Name == name{
			return c.Status(400).JSON(fiber.Map{
				"message": "You already liked this product",
			})
		}
	}
	DB.Model(&User).Association("LikedProducts").Append(&Product)
	fmt.Println("Product Likes: ", Product.UserLikes)

	Product.LikesCount = Product.LikesCount + 1
	fmt.Println("PRODUCT LIKES: ", Product.LikesCount)

	DB.Save(&Product)
	return c.Status(200).JSON(fiber.Map{
		"message": "Product liked",
	})
	
}