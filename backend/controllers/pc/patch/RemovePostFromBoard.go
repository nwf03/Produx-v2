package patch

import (
	"github.com/gofiber/fiber/v2"
  "strings"
  "tutorial/db"
	"github.com/golang-jwt/jwt/v4"
)


func isPrimaryType(t string) bool {
  switch t{
case "bugs":
case "suggestions":
case "announcements":
  return true 
}
  return false
} 
func RemovePostFromBoard(c *fiber.Ctx) error {
	productId := c.Params("productId")
	postId := c.Params("postId")
	field := strings.ToLower(c.Params("field"))
  if !db.ValidType(field){
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

  if isPrimaryType(field) {
    return c.SendStatus(fiber.StatusBadRequest)
  }
  post.RemoveType(field)

  
	return c.Status(200).JSON(post)
}
