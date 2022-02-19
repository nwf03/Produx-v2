package delete

import (
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//var DB = controllers.DB

func DeletePost(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	if !db.ValidType(field) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid field type",
		})
	}
	postId := c.Params("post_id")
	if !mw.ValidFields(field) {
		return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	var product db.Product
	db.DB.First(&product, "name = ?", name)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
	}

	var post db.DeletePost
	switch field {
	//case "Suggestions":
	//	var suggestion db.Suggestion
	//	db.DB.First(&suggestion, "id = ? and user_id = ?", postId, int64(id))
	//	post = &suggestion
	//case "Bugs":
	//	var bug db.Bug
	//	db.DB.First(&bug, "id = ? and user_id = ?", postId, int64(id))
	//	post = &bug
	//case "Announcements":
	//	var announcement db.Announcement
	//	db.DB.First(&announcement, "id = ? and user_id = ?", postId, int64(id))
	//	post = &announcement
	case "Changelogs":
		var changelog db.Changelog
		db.DB.First(&changelog, "id = ? and user_id = ?", postId, int64(id))
		post = &changelog
	default:
		var p db.Post
		db.DB.First(&p, "id = ? and user_id = ? and type=?", postId, int64(id), field)
		post = &p
	}

	if post.GetID() == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Post not found"})
	}
	post.Delete()
	return c.Status(200).JSON(fiber.Map{"message": "Post deleted"})
}
