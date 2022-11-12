package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

// var DB = controllers.DB

func GetUserProducts(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		var allUsers []db.User
		db.DB.Preload("Products").Find(&allUsers)
		return c.Status(200).JSON(allUsers)
	}
	var user db.User
	db.DB.Preload("Products").Preload("FollowedProducts").First(&user, "name = ?", name)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	return c.Status(200).JSON(user)
}
