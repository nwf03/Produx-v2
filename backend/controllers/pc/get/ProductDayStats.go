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
	bugCount, err := db.DB.ProductFieldPostCount(productIdInt, "bug")
	suggestionCount, err := db.DB.ProductFieldPostCount(productIdInt, "suggestion")
	announcementCount, err := db.DB.ProductFieldPostCount(productIdInt, "announcement")
	changelogsCount, err := db.DB.ProductFieldPostCount(productIdInt, "changelog")
	return c.JSON(fiber.Map{
		"bugs":          bugCount,
		"announcements": announcementCount,
		"changelogs":    changelogsCount,
		"suggestions":   suggestionCount,
	})

}
