package get
import (
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
)

type Res struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Products db.Product `json:"products"`
	Bugs []db.Bug `json:"bugs"`
}
func GetLatestPosts(c *fiber.Ctx) error {
	var res Res
	db.DB.Table("users").Preload("Bugs").Scan(&res)
	return c.JSON(res)
}