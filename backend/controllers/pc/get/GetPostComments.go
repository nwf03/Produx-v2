package get

import (
	"strconv"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func GetPostComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	lastId := c.Params("lastId")
	postIdInt, err := strconv.ParseUint(postId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid post id",
		})
	}
	lastIdInt, err := strconv.ParseUint(lastId, 10, 32)
	if err != nil {
		lastIdInt = 0
	}

	comments, err := db.DB.GetPostComments(postIdInt, lastIdInt)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if len(comments) > 0 {
		return c.JSON(fiber.Map{
			"comments": comments,
			"hasMore":  db.DB.GetLastPostCommentID(postIdInt) != comments[len(comments)-1].ID,
			"lastId":   comments[len(comments)-1].ID,
		})
	}
	return c.JSON(fiber.Map{
		"comments": comments,
		"hasMore":  false,
		"lastId":   0,
	})
}
