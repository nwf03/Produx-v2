package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RemoveLikes(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var User db.User
	db.DB.Preload("LikedProducts").Where("name = ?", name).First(&User)
	for _, product := range User.LikedProducts {
		err := db.DB.Model(&product).Association("UserLikes").Delete(User)
		if err != nil {
			return err
		}
	}

	return c.JSON(User)

}
