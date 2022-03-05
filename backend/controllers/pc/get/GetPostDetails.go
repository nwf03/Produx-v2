package get

import (
	"fmt"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetPostDetails(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	var post db.Post
	if db.ValidType(field) {
		query := fmt.Sprintf(`id = %s and type && '{"%s"}'`, postId, field)
		db.DB.Order("created_at desc").Preload("Product").Preload("User").Where(query).Find(&post)
		fmt.Println("post types: ", post.Type)
		return c.JSON(post)
	}
	return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
}
