package post

import (
	"fmt"
	"time"
	"tutorial/controllers"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var DB = controllers.DB

type loginInfo struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(c *fiber.Ctx) error {
	userInfo := new(loginInfo)
	if err := c.BodyParser(userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	fmt.Println("username: " + userInfo.Username)
	user := db.User{}
	// todo hash user passwords
	DB.First(&user, "name = ? AND password = ?", userInfo.Username, userInfo.Password)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "incorrect credentials"})
	}

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
	// todo also return user info too?
	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}
