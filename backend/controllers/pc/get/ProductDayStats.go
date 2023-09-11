package get

import (
	"strconv"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func ProductDayStats(c *fiber.Ctx) error {
	productIdInt, err := strconv.ParseInt(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product id must be an integer",
		})
	}
	bugCount, err := db.DB.ProductFieldPostCount(productIdInt, "bugs", true)
	suggestionCount, err := db.DB.ProductFieldPostCount(productIdInt, "suggestions", true)
	announcementCount, err := db.DB.ProductFieldPostCount(productIdInt, "announcements", true)
	underReviewCount, err := db.DB.ProductFieldPostCount(productIdInt, "under-review", false)
	workingOnCount, err := db.DB.ProductFieldPostCount(productIdInt, "working-on", false)
	doneCount, err := db.DB.ProductFieldPostCount(productIdInt, "done", false)
	return c.JSON(fiber.Map{
		"bugs":          bugCount,
		"announcements": announcementCount,
		"suggestions":   suggestionCount,
		"Under Review":  underReviewCount,
		"Working on ":   workingOnCount,
		"Done":          doneCount,
	})
}
