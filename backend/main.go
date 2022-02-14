package main

import (
	"context"
	"encoding/json"
	"fmt"
	"tutorial/controllers/pc"
	"tutorial/routes"

	"github.com/go-redis/redis/v8"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // use default Addr
	Password: "",               // no password set
	DB:       0,
})
var ctx = context.Background()

func main() {
	var subscriber *redis.PubSub
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
		subscriber = redisClient.Subscribe(ctx, c.Params("id"))
		go func() {
			for {
				msg, err := subscriber.ReceiveMessage(ctx)
				if err != nil {
					panic(err)
				}
				var message pc.Message
				if err = json.Unmarshal([]byte(msg.Payload), &message); err != nil {
					panic(err)
				}
				pc.SendMessage(message, c.Params("id"))
			}
		}()

		var (
			msg []byte
			err error
		)
		for {

			if _, msg, err = c.ReadMessage(); err != nil {
				pc.HandleDisconnect(pc.UserAccs[c.Params("id")][c])
				break
			}
			pc.WSMsgHandler(c, string(msg))

		}
	}))
	fmt.Println("subscriber not nil")
	err := app.Listen(":8000")
	if err != nil {
		return
	}

}
