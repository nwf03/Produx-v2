package get

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"tutorial/controllers/pc/mw"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

type ProductUsersInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Pfp  string `json:"pfp"`
}

func GetProducts(c *fiber.Ctx) error {
	name := strings.TrimSpace(strings.ReplaceAll(c.Params("name"), "%20", " "))
	field := strings.TrimSpace(strings.Title(strings.ToLower(c.Params("field"))))
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
			db.DB.Select("private, user_id, id,name, images, description, verified").First(&product, "name = ?", name)
			if product.ID == 0 {
				return c.Status(404).JSON(fiber.Map{
					"message": "Product not found",
				})
			}
			var Users []ProductUsersInfo
			rows, err := db.DB.Table("users, product_users").Preload("FollowedProducts").Select("users.id, users.name, product_users.role, users.pfp").Where("product_users.product_id = ? and product_users.user_id = users.id", product.ID).Rows()
			if err != nil {
				log.Fatal(err)
			}
			defer func(rows *sql.Rows) {
				err := rows.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(rows)
			for rows.Next() {
				var Pu ProductUsersInfo
				err := rows.Scan(&Pu.ID, &Pu.Name, &Pu.Role, &Pu.Pfp)
				if err != nil {
					log.Fatal(err)
				}
				Users = append(Users, Pu)
			}
			return c.JSON(fiber.Map{
				"product": product,
				"users":   Users,
			})

		}
	}
	var count int64
	if name == "" || name == "-1" && field == "-1" {

		db.DB.Model(&db.Product{}).Find(&products).Count(&count)
		db.DB.Preload("Users").Preload("UserLikes").Preload("PostsCount").Limit(10).Offset(int(pageNum)*10-10).Select("ID", "CreatedAt", "UpdatedAt", "UserID", "name", "description", "images", "likes_count", "posts_count").Find(&products)
		pageCount := mw.GetPageCount(count)
		return c.Status(200).JSON(fiber.Map{"products": products, "pages": pageCount})
	} else if field == "" || field == "-1" {
		query := fmt.Sprintf("SELECT * FROM products WHERE tsv @@ to_tsquery('%s:*');", strings.ReplaceAll(name, " ", "+"))
		rows, err := db.DB.Raw(query).Rows()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				panic(err)
			}
		}(rows)
		for rows.Next() {
			var product db.Product
			err := db.DB.ScanRows(rows, &product)
			if err != nil {
				return err
			}
			products = append(products, product)
		}
		if len(products) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
		}
		pageCount := mw.GetPageCount(int64(len(products)))
		return c.Status(200).JSON(fiber.Map{"products": products, "pages": pageCount})
	}
	return c.Status(404).JSON(fiber.Map{"message": "Product not found"})
}
