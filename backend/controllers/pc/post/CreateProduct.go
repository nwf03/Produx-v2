package post

import (
	"fmt"
	"strings"
	"tutorial/controllers"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var url = "http://localhost:8000/"

var DB = controllers.DB


func CreateProduct(c *fiber.Ctx) error {
	product := new(db.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	product.Name = strings.ToLower(product.Name)
	DB.Where("name = ?", product.Name).First(&product)
	if product.ID != 0 {
		return c.Status(400).JSON(fiber.Map{"error": "product already exists"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	username := claims["name"].(string)
	for i := 1; i < 6; i++ {
		file, err := c.FormFile(fmt.Sprint("image", i))
		if err != nil {
			break
		}
		err = c.SaveFile(file, fmt.Sprintf("./public/%s/%s", username, file.Filename))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		product.Images = append(product.Images, fmt.Sprint(url+"public/"+username+"/"+file.Filename))
	}
	product.UserID = uint(userID)

	fmt.Println("userID: ", userID)
	DB.Create(&product)
	if product.ID == 0 {
		return c.Status(500).JSON(fiber.Map{"error": "product not created"})
	}
	return c.Status(200).JSON(product)
}
