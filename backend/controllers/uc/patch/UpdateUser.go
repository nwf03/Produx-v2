package patch

import (
	"fmt"
	"os"
	"tutorial/db"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UpdateUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)

	var userInfo db.User
	db.DB.Preload("Products").First(&userInfo, "id = ?", uint(userId))
	if userInfo.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	oldName := userInfo.Name
	var pfpUrl string
	pfp, err := c.FormFile("pfp")
	if err == nil {
		err = c.SaveFile(pfp, fmt.Sprintf("./public/%s/%s", userInfo.Name, pfp.Filename))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Error Saving file"})
		}
		pfpUrl = os.Getenv("IMAGE_DOMAIN") + "/public/" + userInfo.Name + "/" + pfp.Filename
	}
	os.Rename(fmt.Sprintf("./public/%s/", oldName), fmt.Sprintf("./public/%s/", userInfo.Name))
	userInfo.Update(c.FormValue("name"), c.FormValue("email"), c.FormValue("password"), pfpUrl)
	db.DB.Save(&userInfo)

	return c.SendStatus(200)
}
