package get

import (
	"strconv"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	afterId := c.Params("afterId")
	afterIdInt, err := strconv.ParseInt(afterId, 0, 64)
	if err != nil {
		afterIdInt = 0
	}

	field := strings.ToLower(c.Params("field"))
	productId := c.Params("productId")
	productIdInt, err := strconv.ParseInt(productId, 0, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "productId must be a number"})
	}
	if !db.ValidType(field) {
		return c.JSON(fiber.Map{"message": "invalid post type"})
	}
	posts, err := db.DB.GetPosts(field, productIdInt, afterIdInt, true)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error"})
	}
	if len(posts) == 0 {
		return c.JSON(fiber.Map{"lastId": 0, "posts": posts, "hasMore": false})
	}
	lastPost := db.DB.GetOldestPost(productIdInt, field)
	return c.JSON(fiber.Map{"lastId": posts[len(posts)-1].GetID(), "posts": posts, "hasMore": lastPost.ID != posts[len(posts)-1].ID})
	// }

}
