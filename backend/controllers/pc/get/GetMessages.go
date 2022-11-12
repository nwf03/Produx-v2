package get

import (
	"strconv"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetMessages(c *fiber.Ctx) error {
	productId := c.Params("productId")
	lastId := c.Params("afterId")
	productIdInt, err := strconv.ParseUint(productId, 10, 32)
	if err != nil {
		productIdInt = 0
	}
	lastIdInt, err := strconv.ParseUint(lastId, 10, 32)
	if err != nil {
		lastIdInt = 0
	}
	messages := db.DB.GetChatMessages(productIdInt, lastIdInt)
	var lastMsgId uint
	if len(messages) > 0 {
		lastMsgId = messages[len(messages)-1].ID
	}
	oldestId := db.DB.GetLastProductMessageId(productIdInt)
	return c.JSON(fiber.Map{
		"messages": messages,
		"lastId":   lastMsgId,
		"hasMore":  oldestId != lastMsgId,
	})
}
