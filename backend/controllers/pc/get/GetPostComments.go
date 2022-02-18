package get

import (
	"errors"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPostComments(c *fiber.Ctx) error {
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	switch field {
	case "bugs":
		var post db.Bug
		err := db.DB.Preload("Product").Preload("Comments").Preload("Comments.User").Preload("User").Where("id = ?", postId).Find(&post).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		return c.JSON(post)
	case "suggestions":
		var post db.Suggestion
		err := db.DB.Preload("Product").Preload("Comments").Preload("Comments.User").Preload("User").Where("id = ?", postId).Find(&post).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		return c.JSON(post)
	case "announcements":
		var post db.Announcement
		err := db.DB.Preload("Product").Preload("Comments").Preload("Comments.User").Preload("User").Where("id = ?", postId).Find(&post).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		return c.JSON(post)
	}
	return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
}
