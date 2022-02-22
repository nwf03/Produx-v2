
package get
import (
	"database/sql"
	"errors"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func GetProduct(c *fiber.Ctx) error {
	name := c.Params("product_name")
	var product db.Product
	var User db.User

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(float64)
	db.DB.First(&User, int(userID))
	err := db.DB.Select("private, user_id, id, name, images, description, verified").Where("name = ?", name).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	isOwner := User.ID == product.UserID
	var Users []ProductUsersInfo
	rows, err := db.DB.Table("users, product_users").Preload("FollowedProduct").Select("users.id, users.name, product_users.role, users.pfp").Where("product_users.product_id = ? and product_users.user_id = users.id", product.ID).Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	for rows.Next() {
		var PUser ProductUsersInfo
		err := rows.Scan(&PUser.ID, &PUser.Name, &PUser.Role, &PUser.Pfp)
		if err != nil {
			panic(err)
		}
		Users = append(Users, PUser)
	}
  var UsersCount int64
  db.DB.Table("product_users").Where("product_id = ?", product.ID).Count(&UsersCount)
  var PostsCount int64
  db.DB.Model(db.Post{}).Where("product_id = ?", product.ID).Count(&PostsCount)
	return c.JSON(fiber.Map{
		"product": product,
		"users":   Users,
		"owner":   isOwner,
    "usersCount": UsersCount,
    "postsCount": PostsCount,
	})

}
