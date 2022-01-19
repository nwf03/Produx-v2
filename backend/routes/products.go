package routes

import (
	"tutorial/controllers/pc/delete"
	"tutorial/controllers/pc/get"
	"tutorial/controllers/pc/patch"
	"tutorial/controllers/pc/post"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Products(app *fiber.App) {
	products := app.Group("/products")
	products.Get("search/:name?/:field?/:page?", get.GetProducts)
	products.Get("/today_top_products", get.TopProducts)
	products.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	products.Post("verify/:name", post.VerifyProduct)
	products.Get("/latest_posts/:field", get.LatestProductPosts)

	products.Post("/create", post.CreateProduct)

	products.Post("/follow", post.FollowProduct)
	products.Post("/unfollow", post.UnfollowProduct)
	products.Post("like/", post.LikeProduct)
	products.Post("dislike/", post.DislikeProduct)

	products.Patch("/update", patch.UpdateProduct)

	products.Post("create/post/:name/:field", post.CreatePost)

	products.Patch("update/post/:name/:field/:post_id", patch.UpdatePost)

	products.Delete("delete/post/:name/:field/:post_id", delete.DeletePost)

}
