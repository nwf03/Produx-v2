package patch

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DislikePost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var User db.User
	db.DB.Where("id = ?", userId).Find(&User)
	post := new(db.Post)

	db.DB.First(&post, "id = ?", postId)
	err := post.Dislike(User)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
