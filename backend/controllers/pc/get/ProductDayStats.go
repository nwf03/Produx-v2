package get

import (
	"strconv"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func ProductDayStats(c *fiber.Ctx) error {
	// convert string to uint
	productIdInt, err := strconv.ParseInt(c.Params("productId"), 10, 32)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product id must be an integer",
		})
	}
	bugCount, err := db.DB.ProductFieldPostCount(productIdInt, "bugs")
	suggestionCount, err := db.DB.ProductFieldPostCount(productIdInt, "suggestions")
	announcementCount, err := db.DB.ProductFieldPostCount(productIdInt, "announcements")
	underReviewCount, err := db.DB.ProductFieldPostCount(productIdInt, "under-review")
	workingOnCount, err := db.DB.ProductFieldPostCount(productIdInt, "working-on")
	doneCount, err := db.DB.ProductFieldPostCount(productIdInt, "done")
	return c.JSON(fiber.Map{
		"bugs":          bugCount,
		"announcements": announcementCount,
		"suggestions":   suggestionCount,
		"underReview":   underReviewCount,
		"workingOn":     workingOnCount,
		"done":          doneCount,
	})
}
