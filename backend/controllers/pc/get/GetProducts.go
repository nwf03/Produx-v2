package get

import (
	"fmt"
	"strconv"
	"strings"
	"tutorial/controllers"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

var DB = controllers.DB


func GetProducts(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
	fmt.Println("filed: " + field)
	var products []db.Product

	page := c.Params("page")
	var pageNum int64 = 1
	var err error
	if page != "" {
		pageNum, err = strconv.ParseInt(page, 0, 8)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "page must be a number"})
		}
		if pageNum == 0 {
			var product db.Product
			DB.Preload("Users").First(&product, "name = ?", name)
			if product.ID == 0 {
				return c.Status(404).JSON(fiber.Map{
					"error": "Product not found",
				})
			}
			return c.JSON(product)

		}
	}
	var count int64
	if name == "" || name == "-1" && field == "-1" {

		DB.Model(&db.Product{}).Find(&products).Count(&count)
		DB.Preload("Users").Preload("UserLikes").Preload("PostsCount").Limit(10).Offset(int(pageNum)*10-10).Select("ID", "CreatedAt", "UpdatedAt", "UserID", "name", "description", "images", "likes_count", "posts_count").Find(&products)
		pageCount := mw.GetPageCount(count)
		return c.Status(200).JSON(fiber.Map{"products": products, "pages": pageCount})
	} else if field == "" || field == "-1" {

		DB.Model(&db.Product{}).Find(&products, "name LIKE  ?", "%"+name+"%").Count(&count)
		DB.Preload("Users").Preload("UserLikes").Preload("PostsCount").Limit(10).Offset(int(pageNum)*10-10).Select("ID", "CreatedAt", "UpdatedAt", "UserID", "name", "description", "images", "likes_count", "posts_count").Find(&products, "name LIKE  ?", "%"+name+"%")

		if len(products) == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}
		pageCount := mw.GetPageCount(count)
		return c.Status(200).JSON(fiber.Map{"products": products, "pages": pageCount})
	} else if mw.ValidFields(field) {
		var product db.Product
		DB.Preload(field).Preload(field+".User").Where("name = ?", name).First(&product)

		if product.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}
		switch field {
		case "Suggestions":
			var suggestions []db.Suggestion
			DB.Model(&db.Suggestion{}).Count(&count)
			DB.Preload("User").Limit(10).Offset(int(pageNum)*10-10).Order("created_at desc").Find(&suggestions, "product_id = ?", product.ID)
			return c.Status(200).JSON(fiber.Map{"posts": suggestions, "pages": mw.GetPageCount(count)})
		case "Bugs":
			var bugs []db.Bug
			DB.Model(&db.Bug{}).Count(&count)

			DB.Preload("User").Limit(10).Offset(int(pageNum)*10-10).Order("created_at desc").Find(&bugs, "product_id = ?", product.ID)
			return c.Status(200).JSON(fiber.Map{"pages": mw.GetPageCount(count), "posts": bugs})
		case "Changelogs":
			DB.Model(&db.Changelog{}).Count(&count)
			var changelogs []db.Changelog
			DB.Preload("User").Limit(10).Offset(int(pageNum)*10-10).Order("created_at desc").Find(&changelogs, "product_id = ?", product.ID)
			return c.Status(200).JSON(fiber.Map{"posts": changelogs, "pages": mw.GetPageCount(count)})
		case "Announcements":
			var announcements []db.Announcement
			DB.Model(&db.Announcement{}).Count(&count)
			DB.Preload("User").Limit(10).Offset(int(pageNum)*10-10).Order("created_at desc").Find(&announcements, "product_id = ?", product.ID)
			return c.Status(200).JSON(fiber.Map{"posts": announcements, "pages": mw.GetPageCount(count)})
		}
		return c.Status(404).JSON(fiber.Map{"error": "Field not found"})
	} else {
		return c.Status(400).JSON(fiber.Map{"error": "invalid field"})
	}
}
