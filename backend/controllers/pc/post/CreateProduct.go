package post

import (
	"fmt"
	"os"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var url = "http://localhost:8000/"

func CreateProduct(c *fiber.Ctx) error {
	product := new(db.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	product.Name = strings.ToLower(product.Name)
	db.DB.Where("name = ?", product.Name).First(&product)
	if product.ID != 0 {
		return c.Status(400).JSON(fiber.Map{"message": "product already exists"})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	var userInfo db.User
	db.DB.First(&userInfo, "id = ?", userID)
	for i := 1; i < 6; i++ {
		file, err := c.FormFile(fmt.Sprint("image", i))
		if err != nil {
			break
		} else if i == 1 {
			os.Mkdir("./public/"+userInfo.Name+"/"+product.Name, 0755)
		}
		imgUrl := fmt.Sprintf("./public/%s/%s/%s", userInfo.Name, product.Name, file.Filename)

		err = c.SaveFile(file, imgUrl)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		product.Images = append(product.Images, url+imgUrl[2:])
	}
	product.UserID = uint(userID)

	db.DB.Create(&product)
	if product.ID == 0 {
		return c.Status(500).JSON(fiber.Map{"message": "product not created"})
	}
	relationship := db.ProductUser{
		UserID:    uint(userID),
		ProductID: product.ID,
		Role:      "Owner",
	}
	db.DB.Save(&relationship)
	return c.Status(200).JSON(product)
}
