package get

import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func GetFromPostsBoard(c *fiber.Ctx) error {
	productId := c.Params("productId")
	var workingOn []db.Post
	var done []db.Post
	var underReview []db.Post
	db.DB.Find(&workingOn, "product_id = ? and type && ?", productId, pq.StringArray{"working-on"})
	db.DB.Find(&done, "product_id = ? and type && ?", productId, pq.StringArray{"done"})
	db.DB.Find(&underReview, "product_id = ? and type && ?", productId, pq.StringArray{"under-review"})
	return c.JSON(fiber.Map{
		"workingOn":   workingOn,
		"underReview": underReview,
		"done":        done,
	})
}
