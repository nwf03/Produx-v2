package get

import (
	"strings"
	"tutorial/db"
  "fmt"
	"github.com/gofiber/fiber/v2"
)

func GetPostComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	var post db.Post
	if db.ValidType(field) {
    query := fmt.Sprintf(`id = %s and type && '{"%s"}'`, postId, field)
    db.DB.Preload("Product").Preload("Comments").Preload("Comments.User").Preload("User").Where(query).Find(&post)
    fmt.Println("post types: ", post.Type)
		return c.JSON(post)
	}
	return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
}
