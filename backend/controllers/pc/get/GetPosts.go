package get

import (
	"fmt"
	"strconv"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        User   `json:"user"`
	UserId      uint   `json:"userID"`
	//UserID    int `json:"user_id"`
	//ProductID int `json:"product_id"`
}
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pfp   string `json:"pfp"`
}

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
	if !(field == "changelogs") && !db.ValidType(field) {
		return c.JSON(fiber.Map{"message": "invalid post type"})
	}
	switch field {
	//case "suggestions":
	//	suggestions, err := db.DB.GetSuggestions(productIdInt, afterIdInt, true)
	//	if err != nil {
	//		return c.Status(500).JSON(fiber.Map{"message": "error"})
	//	}
	//	if len(suggestions) == 0 {
	//		return c.Status(404).JSON(fiber.Map{"message": "no suggestions"})
	//	}
	//	lastId := suggestions[len(suggestions)-1].GetID()
	//	oldestSug := db.DB.GetOldestSuggestions(productIdInt)
	//	return c.JSON(fiber.Map{"lastId": lastId, "posts": suggestions, "hasMore": oldestSug.ID != lastId})
	//case "bugs":
	//	bugs, err := db.DB.GetBugs(productIdInt, afterIdInt, true)
	//	if err != nil {
	//		fmt.Println(err)
	//		return c.Status(500).JSON(fiber.Map{"message": "error"})
	//	}
	//	if len(bugs) == 0 {
	//		return c.Status(404).JSON(fiber.Map{"message": "no bugs"})
	//	}
	//	oldestBug := db.DB.GetOldestBugs(productIdInt)
	//	lastId := bugs[len(bugs)-1].GetID()
	//	fmt.Println("Last id: ", lastId)
	//	fmt.Println("oldest id: ", oldestBug.ID)
	//	return c.JSON(fiber.Map{"lastId": lastId, "posts": bugs, "hasMore": oldestBug.ID != lastId})
	case "changelogs":
		changelogs, err := db.DB.GetChangelogs(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(changelogs) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no changelogs"})
		}
		return c.JSON(fiber.Map{"lastId": changelogs[0].GetID(), "posts": changelogs})

	default:
		posts, err := db.DB.GetPosts(field, productIdInt, afterIdInt, true)
		fmt.Println("error while getting posts: ", err)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(posts) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no posts"})
		}
		return c.JSON(fiber.Map{"lastId": posts[0].GetID(), "posts": posts})
		//case "announcements":
		//	announcements, err := db.DB.GetAnnouncements(productIdInt, afterIdInt, true)
		//	if err != nil {
		//		return c.Status(500).JSON(fiber.Map{"message": "error"})
		//	}
		//	if len(announcements) == 0 {
		//		return c.Status(404).JSON(fiber.Map{"message": "no announcements"})
		//	}
		//	return c.JSON(fiber.Map{"lastId": announcements[0].GetID(), "posts": announcements})

	}

}
