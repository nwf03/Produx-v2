package get

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"tutorial/db"
)

func IsProductFollowed(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	var User db.ProductUser
	productId := c.Params("product_id")
	//convert product id to uint64
	var Product db.Product
	if err := db.DB.Select("user_id").Where("id = ?", productId).First(&Product).Error; err != nil {
		return c.Status(500).JSON(err)
	}
	if uint(userID) == Product.UserID {
		return c.Status(200).JSON(fiber.Map{
			"followed": true,
		})
	}
	err := db.DB.Model(&db.ProductUser{}).Where("user_id = ? AND product_id = ?", userID, productId).First(&User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(200).JSON(fiber.Map{"followed": false})
	}
	return c.Status(200).JSON(fiber.Map{"followed": true})

}
