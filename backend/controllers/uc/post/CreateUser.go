package post

import (
	"fmt"
	"log"
	"os"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var url = "http://localhost:8000/"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Pfp      string `json:"pfp"`
}

func CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return err
	}
	var existingUser db.User
	db.DB.First(&existingUser, "email = ?", req.Email)
	if existingUser.Email != "" {
		return c.Status(400).JSON(fiber.Map{"message": "User already exists"})
	}
	file, err := c.FormFile("pfp")
	if err == nil {
		err := os.Mkdir("./public/"+req.Name, 0755)
		if err != nil {
			log.Fatal(err)
		}
		err = c.SaveFile(file, fmt.Sprintf("./public/%s/%s", req.Name, file.Filename))
		if err != nil {
			return err
		}
		req.Pfp = url + "public/" + req.Name + "/" + file.Filename
	}

	if req.Name == "" || req.Password == "" || req.Email == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Please fill all the fields"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &db.User{
		Name:     req.Name,
		Email:    req.Email,
		Pfp:      req.Pfp,
		Password: string(hash),
	}

	db.DB.Create(user)

	tokenString, err := createJWTToken(user.ID, 72)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}
