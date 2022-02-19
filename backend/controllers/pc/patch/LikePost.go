package patch

import (
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LikePost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	if !db.ValidType(field) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid field",
		})
	}
	productId := c.Params("productId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var User db.User
	db.DB.Where("id = ?", userId).Find(&User)
	var post db.PostLiker

	switch field {
	case "bugs":
	case "suggestions":
	case "announcements":
		post = new(db.Post)
	case "changelogs":
		post = new(db.Changelog)
	}
	db.DB.First(&post, "id = ? AND product_id = ? and type=?", postId, productId, field)
	err := post.Like(User)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
