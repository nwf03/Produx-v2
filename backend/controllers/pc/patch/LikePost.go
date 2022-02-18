package patch

import (
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LikePost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := c.Params("field")
	productId := c.Params("productId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var User db.User
	db.DB.Where("id = ?", userId).Find(&User)
	var post db.PostLiker

	field = strings.ToLower(field)
	switch field {
	case "bugs":
		post = new(db.Bug)
	case "suggestions":
		post = new(db.Suggestion)
	case "announcements":
		post = new(db.Announcement)
	case "changelogs":
		post = new(db.Changelog)
	}
	db.DB.First(&post, "id = ? AND product_id = ?", postId, productId)
	err := post.Like(User)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{
    "message": "success",
  })
}
