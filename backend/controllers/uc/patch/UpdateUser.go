package patch

import (
	"time"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UpdateUser(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var userInfo db.User
	db.DB.First(&userInfo, "name = ?", name)
	if userInfo.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	updatedUsername := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	if updatedUsername == "" && email == "" && password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "username, email, or password are required"})
	}
	if updatedUsername != "" {
		userInfo.Name = updatedUsername
	}
	if password != "" {
		userInfo.Password = password
	}
	if email != "" {
		userInfo.Email = email
	}

	claims2 := jwt.MapClaims{
		"name": userInfo.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	db.DB.Save(&userInfo)

	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}
