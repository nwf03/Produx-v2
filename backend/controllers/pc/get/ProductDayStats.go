package get

import (
	"time"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func ProductDayStats(c *fiber.Ctx) error {
	var count int64 
	today := time.Now()
	year, month, day := today.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, today.Location())
	dayEnd := time.Date(year, month, day, 23, 59, 59, 0, today.Location())
	db.DB.Model(db.Bug{}).Where("product_id = ? and created_at BETWEEN ? AND ?", c.Params("productId"),dayStart, dayEnd).Count(&count)
	return c.JSON(fiber.Map{"count": count})
}