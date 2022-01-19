package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func RemoveLikes(c *fiber.Ctx) error{
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var User db.User
	DB.Preload("LikedProducts").Where("name = ?", name).First(&User)
	for _, product := range User.LikedProducts{
		DB.Model(&product).Association("UserLikes").Delete(User)
	}
	
	return c.JSON(User)

}