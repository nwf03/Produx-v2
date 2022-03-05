package get

import (
	"strconv"
	"strings"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LatestProductPosts(c *fiber.Ctx) error {
	field := strings.TrimSpace(strings.ToLower(c.Params("field")))
	lastId := c.Params("lastId")
	lastIdInt, err := strconv.ParseUint(lastId, 10, 32)
	if err != nil {
		lastIdInt = 0
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	var User db.User
	db.DB.Preload("FollowedProducts").Where("id = ?", uint(id)).First(&User)

	if len(User.FollowedProducts) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "no products found"})
	}

	posts, err := db.DB.GetFollowedProductsPosts(User, field, lastIdInt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid field"})
	}
	oldestPostId := db.DB.GetOldestFollowedProductsPost(User.FollowedProducts, field, lastIdInt)
	if len(posts) == 0 {
		return c.JSON(fiber.Map{
			"posts":   posts,
			"lastId":  0,
			"hasMore": false,
		})
	}
	lastPostId := posts[len(posts)-1].ID

	return c.JSON(fiber.Map{
		"posts":   posts,
		"lastId":  lastId,
		"hasMore": oldestPostId != lastPostId,
	})
}
