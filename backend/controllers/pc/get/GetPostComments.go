package get

import (
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetPostComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	var post Post
	if db.ValidType(field) {
		db.DB.Preload("Product").Preload("Comments").Preload("Comments.User").Preload("User").Where("id = ? and type = ?", postId, field).Find(&post)
		return c.JSON(post)
	}
	return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
}
