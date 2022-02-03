package get

import (
	"fmt"
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LatestProductPosts(c *fiber.Ctx) error {
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	if !mw.ValidFields(field) {
		return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	var User db.User
	q1 := fmt.Sprint("FollowedProducts.", field)
	q2 := fmt.Sprint(q1, ".User")
	db.DB.Preload("FollowedProducts").Preload(q1).Preload(q2).Where("id = ?", id).First(&User)

	if (len(User.FollowedProducts) == 0) {
		return c.Status(404).JSON(fiber.Map{"message": "no products found"})
	}
	
	fmt.Println(field)
	fmt.Println(User)
	var productIds []uint
	for _, product := range User.FollowedProducts {
		productIds = append(productIds, product.ID)
	}

	switch field{
		case "Suggestions":
			var Suggestions []db.Suggestion
			db.DB.Preload("Product").Preload("User").Where("product_id IN (?)", productIds).Order("created_at desc").Find(&Suggestions)
			return c.JSON(Suggestions)
		case "Bugs":
			var Bugs []db.Bug
			db.DB.Preload("Product").Preload("User").Where("product_id IN (?)", productIds).Order("created_at desc").Find(&Bugs)
			return c.JSON(Bugs)
		case "Changelogs":
			var Changelogs []db.Changelog
			db.DB.Preload("Product").Preload("User").Where("product_id IN (?)", productIds).Order("created_at desc").Find(&Changelogs)
			return c.JSON(Changelogs)
		case "Announcements":
			var Announcements []db.Announcement
			db.DB.Preload("Product").Preload("User").Where("product_id IN (?)", productIds).Order("created_at desc").Find(&Announcements)
			return c.JSON(Announcements)
		
	}
	return c.JSON(User.FollowedProducts)
	


}