package delete

import (
	"fmt"
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DeletePost(c *fiber.Ctx) error {
	productId := c.Params("productId")
	field := strings.TrimSpace(strings.ToLower(c.Params("field")))
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
	db.DB.First(&product, "id = ?", productId)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
	}

	var post db.DeletePost
	var p db.Post
	query := fmt.Sprintf(`id = %s and user_id = %d and type && '{"%s"}'`, postId, uint(id), field)
	db.DB.First(&p, query)
	post = &p

	if post.GetID() == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Post not found"})
	}
	post.Delete()
	return c.Status(200).JSON(fiber.Map{"message": "Post deleted"})
}
