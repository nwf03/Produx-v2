package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var userInfo db.User
	DB.Preload("FollowedProducts").Preload("LikedProducts").Preload("Products").First(&userInfo, "name = ?", name)
	if userInfo.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	return c.Status(200).JSON(userInfo)
}
