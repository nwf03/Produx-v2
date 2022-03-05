package post

import (
	"fmt"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type FPPReq struct {
	AccessToken string `json:"accessToken"`
}

func FollowProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var User db.User
	db.DB.Where("id = ?", userId).First(&User)
	productName := strings.TrimSpace(strings.ReplaceAll(c.Params("product_name"), "%20", " "))

	var Product db.Product
	db.DB.Where("name = ?", productName).First(&Product)
	if !Product.Private || uint(userId) == Product.UserID {
		err := Product.Follow(User.ID)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(fiber.Map{"status": "success"})
	}

	var req FPPReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(503).JSON(fiber.Map{"status": "error", "message": "provide access token"})
	}
	if req.AccessToken != Product.AccessToken {
		fmt.Println(req.AccessToken)
		fmt.Println(Product.AccessToken)
		return c.Status(503).JSON(fiber.Map{"message": "AccessToken is not correct"})
	}

	err := Product.Follow(User.ID)
	if err != nil {
		return c.Status(503).SendString("Error: " + err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Followed",
	})
}
