package post

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"tutorial/db"
)

type CCReq struct {
	Field   string `json:"field"`
	PostID  string `json:"postId"`
	Comment string `json:"comment"`
}

func CreateComment(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var req CCReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	var UserInfo db.User
	err := db.DB.Where("id = ?", userId).First(&UserInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.ErrUnauthorized
	}
	switch req.Field {
	case "bugs":
		var BugInfo db.Bug
		err := db.DB.Where("id = ?", req.PostID).First(&BugInfo).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		var CommentInfo db.BugComment
		CommentInfo.UserID = UserInfo.ID
		CommentInfo.BugID = BugInfo.ID
		CommentInfo.User = UserInfo
		CommentInfo.Comment = req.Comment
		err = db.DB.Create(&CommentInfo).Error
		if err != nil {
			return err
		}
		return c.Status(200).JSON(CommentInfo)
	case "suggestion":
		var SuggestionInfo db.Suggestion
		err := db.DB.Where("id = ?", req.PostID).First(&SuggestionInfo).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		var CommentInfo db.SuggestionComment
		CommentInfo.UserID = UserInfo.ID
		CommentInfo.SuggestionID = SuggestionInfo.ID
		CommentInfo.User = UserInfo
		CommentInfo.Comment = req.Comment
		err = db.DB.Create(&CommentInfo).Error
		if err != nil {
			return err
		}
		return c.Status(200).JSON(CommentInfo)
	}
	return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
}
