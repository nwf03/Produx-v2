package main

import (
	"tutorial/controllers/pc"
	"tutorial/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
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
	app.Use("/ws", func(c *fiber.Ctx) error { // IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			msg []byte
			err error
		)
		for {

			if _, msg, err = c.ReadMessage(); err != nil {
				pc.HandleDisconnect(c)
				break
			}
			pc.WSMsgHandler(c, string(msg))

		}

	}))
	err := app.Listen(":8000")
	if err != nil {
		return
	}

}
