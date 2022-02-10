package get

import (
	"fmt"
	"strconv"
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

	field := c.Params("field")
	productId := c.Params("productId")
	productIdInt, err := strconv.ParseInt(productId, 0, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "productId must be a number"})
	}

	switch field {
	case "suggestions":
		suggestions, err := db.DB.GetBugs(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(suggestions) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no suggestions"})
		}
		return c.JSON(fiber.Map{"lastId": suggestions[len(suggestions)-1].GetID(), "posts": suggestions})
	case "bugs":
		bugs, err := db.DB.GetBugs(productIdInt, afterIdInt, true)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(bugs) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no bugs"})
		}
		return c.JSON(fiber.Map{"lastId": bugs[len(bugs)-1].GetID(), "posts": bugs})
	case "changelogs":
		changelogs, err := db.DB.GetChangelogs(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(changelogs) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no changelogs"})
		}
		return c.JSON(fiber.Map{"lastId": changelogs[0].GetID(), "posts": changelogs})

	case "announcements":
		announcements, err := db.DB.GetAnnouncements(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		if len(announcements) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "no announcements"})
		}
		return c.JSON(fiber.Map{"lastId": announcements[0].GetID(), "posts": announcements})

	}

	return c.JSON(fiber.Map{"message": "error"})
}
