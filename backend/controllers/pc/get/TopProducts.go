package get

import (
	"time"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func TopProducts(c *fiber.Ctx) error {
	var Products []db.Product
	today := time.Now()
	year, month, day := today.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, today.Location())
	dayEnd := time.Date(year, month, day, 23, 59, 59, 0, today.Location())
	db.DB.Where("private = false and created_at BETWEEN ? AND ?", dayStart, dayEnd).Order("likes_count desc").Limit(10).Find(&Products)
	return c.JSON(Products)
}
