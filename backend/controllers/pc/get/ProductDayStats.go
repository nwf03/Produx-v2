package get

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tutorial/db"
)

func ProductDayStats(c *fiber.Ctx) error {
	//convert string to uint
	productIdInt, err := strconv.ParseInt(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product id must be an integer",
		})
	}
	bugCount := db.DB.ProductFieldPostCount(productIdInt, db.Bugs)
	suggestionCount := db.DB.ProductFieldPostCount(productIdInt, db.Suggestions)
	announcementCount := db.DB.ProductFieldPostCount(productIdInt, db.Announcements)
	changelogsCount := db.DB.ProductFieldPostCount(productIdInt, db.Changelogs)
	return c.JSON(fiber.Map{
		"bugs":          bugCount,
		"announcements": announcementCount,
		"changelogs":    changelogsCount,
		"suggestions":   suggestionCount,
	})

}
