package post

import (
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

func CreatePost(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.ToLower(c.Params("field")))
	if !mw.ValidFields(field) {
		return c.Status(400).JSON(fiber.Map{"error": "invalid field"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	var product db.Product
	db.DB.First(&product, "name = ?", name)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	newField := new(db.Post)

	var userInfo db.User
	db.DB.First(&userInfo, "id = ?", id)

	//switch field {
	//	newField = new(db.Suggestion)
	//	newField.(*db.Suggestion).UserID = uint(id)
	//	newField.(*db.Suggestion).ProductID = product.ID
	//	newField = new(db.Bug)
	//	newField.(*db.Bug).UserID = uint(id)
	//
	//	newField.(*db.Bug).ProductID = product.ID
	// case "announcements":
	//case "suggestions":
	//case "bugs":
	//	newField.UserID = uint(id)
	//	newField.ProductID = product.ID
	// case "changelogs":
	// 	newField = new(db.Changelog)
	// 	newField.(*db.Changelog).UserID = uint(id)
	// 	newField.(*db.Changelog).ProductID = product.ID
	//}
	if err := c.BodyParser(newField); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	newField.UserID = uint(id)
	newField.ProductID = product.ID
	newField.Type = pq.StringArray{field} 
	db.DB.Create(newField)
	return c.Status(200).JSON(newField)
}
