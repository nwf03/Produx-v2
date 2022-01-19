package main

import (
	"tutorial/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/public", "./public")
	routes.Users(app)

	routes.Products(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	err := app.Listen(":8000")
	if err != nil {
		return
	}

}
