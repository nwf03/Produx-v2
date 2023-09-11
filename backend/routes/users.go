package routes

import (
	"tutorial/controllers/uc/get"
	"tutorial/controllers/uc/patch"
	"tutorial/controllers/uc/post"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Users(app *fiber.App) {
	users := app.Group("/users")

	users.Get("/:name?", get.GetUserProducts)

	users.Post("/create", post.CreateUser)

	users.Post("/login", post.Login)

	users.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	users.Get("/removeLikes", get.RemoveLikes)
	users.Patch("/update/", patch.UpdateUser)
	users.Get("/user/get", get.GetUserInfo)
}
