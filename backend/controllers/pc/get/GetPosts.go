package get

import (
	"fmt"
	"gorm.io/gorm"
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
	gorm.Model
	Name             string          `json:"name"`
	Password         string          `json:"-"`
	Email            string          `json:"email"`
	Pfp              string          `json:"pfp"`
	Products         []db.Product    `json:"products"`
	FollowedProducts []db.Product    `json:"followed_products"`
	Suggestions      []db.Suggestion `json:"suggestions,omitempty"`
	Bugs             []db.Bug        `json:"bugs,omitempty"`
	Changelogs       []db.Changelog  `json:"changeLogs,omitempty"`
	LikedProducts    []db.Product    `json:"liked_products"`

	LikedSuggestions   []db.Suggestion   `json:"liked_suggestions"`
	LikedBugs          []db.Bug          `json:"liked_bugs"`
	LikedChangelogs    []db.Changelog    `json:"liked_changelogs"`
	LikedAnnouncements []db.Announcement `json:"liked_announcements"`

	DislikedAnnouncements []db.Announcement `json:"disliked_announcements"`
	DislikedSuggestions   []db.Suggestion   `json:"disliked_suggestions" gorm:"many2many:disliked_suggestions;"`
	DislikedBugs          []db.Bug          `json:"disliked_bugs" gorm:"many2many:disliked_bugs;"`
	DislikedChangelogs    []db.Changelog    `json:"disliked_changelogs" gorm:"many2many:disliked_changelogs;"`
	Messages              []db.Message      `json:"messages,omitempty"`
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
		//suggestions, err := db.DB.GetBugs(productIdInt, afterIdInt, true)
		var suggestions []Post
		var err error
		db.DB.Preload("User").Model(db.Suggestion{}).Find(&suggestions, "product_id = ? and id > ?", productIdInt, afterIdInt)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		//return c.JSON(fiber.Map{"lastId": suggestions[0].GetID(), "posts": suggestions})
		return c.JSON(suggestions)
	case "bugs":
		bugs, err := db.DB.GetBugs(productIdInt, afterIdInt, true)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		return c.JSON(fiber.Map{"lastId": bugs[0].GetID(), "posts": bugs})
	case "changelogs":
		changelogs, err := db.DB.GetChangelogs(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		return c.JSON(fiber.Map{"lastId": changelogs[0].GetID(), "posts": changelogs})

	case "announcements":
		announcements, err := db.DB.GetAnnouncements(productIdInt, afterIdInt, true)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error"})
		}
		return c.JSON(fiber.Map{"lastId": announcements[0].GetID(), "posts": announcements})

	}

	return c.JSON(fiber.Map{"message": "error"})
}
