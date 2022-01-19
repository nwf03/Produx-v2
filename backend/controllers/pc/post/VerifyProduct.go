package post

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

func VerifyProduct(c *fiber.Ctx) error{
	var Product db.Product
	name := c.Params("name")
	DB.First(&Product, "name = ?", name)
	if Product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
		})
	}
	Product.Verified = true
	DB.Save(&Product)
	return c.Status(200).JSON(Product)

}