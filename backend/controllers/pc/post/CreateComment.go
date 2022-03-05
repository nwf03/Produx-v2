package post

import (
	"errors"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type CCReq struct {
	Field   string `json:"field"`
	PostID  int    `json:"postId"`
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
	req.Field = strings.ToLower(req.Field)
	if !db.ValidType(req.Field) {
		return fiber.ErrBadRequest
	}
	var PostInfo db.Post
	err = db.DB.Where("id = ?", req.PostID).First(&PostInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.ErrNotFound
	}

	err = db.DB.Where("id = ?", req.PostID).First(&PostInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.ErrUnauthorized
	}
	var CommentInfo db.Comment
	CommentInfo.UserID = UserInfo.ID
	CommentInfo.PostID = PostInfo.ID
	CommentInfo.User = UserInfo
	CommentInfo.Comment = req.Comment
	err = db.DB.Create(&CommentInfo).Error
	if err != nil {
		return err
	}
	return c.Status(200).JSON(CommentInfo)
}
