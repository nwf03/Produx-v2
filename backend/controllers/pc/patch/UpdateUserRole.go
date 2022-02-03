package patch

import (
	"fmt"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ReqBody struct {
	Role string `json:"role"`
}

func UpdateUserRole(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	var userR db.ProductUser
	var Product db.Product
	db.DB.First(&Product, c.Params("product_id"))
	fmt.Println("userId", userId)
	fmt.Println("productId", Product.UserID)
	if uint(userId) != Product.UserID {
		return c.Status(401).JSON(fiber.Map{
			"message": "You are not authorized to update this product",
		})
	}
	db.DB.Model(db.ProductUser{}).First(&userR, "user_id = ? and product_id = ?", c.Params("user_id"), Product.ID)
	//	parse body
	var body ReqBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	userR.Role = body.Role
	db.DB.Save(&userR)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Roles",
		"data":    userR,
	})
}
