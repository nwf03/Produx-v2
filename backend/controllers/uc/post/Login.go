package post

import (
	"fmt"
	"time"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// var DB = controllers.DB

type loginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	userInfo := new(loginInfo)
	if err := c.BodyParser(userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := db.User{}
	db.DB.First(&user, "name = ?", userInfo.Username)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "email or password incorrect")
	}

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "incorrect credentials"})
	}

	tokenString, err := createJWTToken(user.ID, 72)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// todo also return user info too?
	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}

func createJWTToken(userID uint, lifespan time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * lifespan).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
