package post

import (
	"fmt"
	"os"
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

	if err := c.BodyParser(newField); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
  postId := uuid.NewString()
  err := os.Mkdir("./public/"+userInfo.Name+"/"+postId, 0755) 
  if err != nil{
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  for i := 1; i <= 3; i++ {
    file, err := c.FormFile(fmt.Sprint("image", i)) 
    if err != nil{
      break
    }
    imageUrl := fmt.Sprintf("public/%s/%s/%s", userInfo.Name, postId,file.Filename)
    savePath := fmt.Sprintf("./%s", imageUrl)
    err = c.SaveFile(file, savePath)
    if err != nil{
      c.SendStatus(fiber.StatusInternalServerError)
    }
    newField.Images = append(newField.Images, url+ imageUrl)
  }

	newField.UserID = uint(id)
	newField.ProductID = product.ID
	newField.Type = pq.StringArray{field} 
	db.DB.Create(newField)
	return c.Status(200).JSON(newField)
}
