package patch

import (
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UpdatePost(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	postId := c.Params("post_id")
	newTitle := c.FormValue("title")
	newDescription := c.FormValue("description")
	newVersion := c.FormValue("version")
	if newTitle == "" && newDescription == "" && newVersion == "" {
		return c.Status(400).JSON(fiber.Map{"message": "title or description is required"})
	}

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

	var post db.Post
	db.DB.First(post, "id = ? and user_id = ? and type=?", postId, int64(id), field)
	if post.GetID() == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Post not found"})
	}
	post.Update(newTitle, newDescription)
	return c.JSON(post)
	// }
}
