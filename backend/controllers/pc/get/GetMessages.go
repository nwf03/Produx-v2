package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetMessages(c *fiber.Ctx) error {
	productId := c.Params("productId")
	lastId := c.Params("afterId")
	if lastId == "" {
		lastId = "0"
	}
	var Messages []db.Message
	db.DB.Preload("User").Where("product_id = ? and id > ?", productId, lastId).Limit(15).Find(&Messages)
	return c.JSON(Messages)
}
