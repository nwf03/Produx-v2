package patch

import (
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

func hasSecondaryTypes(types pq.StringArray) (bool, string) {
	for _, t := range types {
		if t == "working-on" || t == "done" || t == "under-review" {
			return true, t
		}
	}
	return false, ""
}

func AddPostToBoard(c *fiber.Ctx) error {
	productId := c.Params("productId")
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
	if !db.ValidType(field) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid type",
		})
	}
	var product db.Product
	var post db.Post
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	db.DB.First(&product, productId)

	if uint(id) != product.UserID {
		return fiber.ErrUnauthorized
	}
	db.DB.First(&post, postId)

	if yes, t := hasSecondaryTypes(post.Type); yes {
		post.RemoveType(t)
	}
	post.AddType(field)

	return c.Status(200).JSON(post)
}
