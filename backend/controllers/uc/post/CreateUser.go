package post

import (
	"fmt"
	"log"
	"os"
	"time"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var url = "http://localhost:8000/"

func CreateUser(c *fiber.Ctx) error {
	user := new(db.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}
	var existingUser db.User
	DB.First(&existingUser, "email = ?", user.Email)
	if existingUser.Email != "" {
		return c.Status(400).JSON(fiber.Map{"message": "User already exists"})
	}
	file, err := c.FormFile("pfp")
	if err == nil {
		err := os.Mkdir("./public/"+user.Name, 0755)
		if err != nil {
			log.Fatal(err)
		}
		err = c.SaveFile(file, fmt.Sprintf("./public/%s/%s", user.Name, file.Filename))
		if err != nil {
			return err
		}
		user.Pfp = url + "public/" + user.Name + "/" + file.Filename
	}

	if user.Name == "" || user.Password == "" || user.Email == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Please fill all the fields"})
	}
	DB.Create(user)

	claims := jwt.MapClaims{
		"name": user.Name,
		"id":   user.ID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}
