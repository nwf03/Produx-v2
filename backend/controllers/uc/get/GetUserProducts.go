package get

import (
	"fmt"
	"tutorial/controllers"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

var DB = controllers.DB


func GetUserProducts(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		var allUsers []db.User
		DB.Preload("Products").Find(&allUsers)
		return c.Status(200).JSON(allUsers)
	}
	var user db.User
	DB.Preload("Products").Preload("FollowedProducts").First(&user, "name = ?", name)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	fmt.Println("user products: ", user.Products)
	return c.Status(200).JSON(user)
}
