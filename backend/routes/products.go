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

	products.Get("/comments/:field/:postId", get.GetPostComments)
	products.Get("/isFollowed/:product_id", get.IsProductFollowed)
	products.Get("messages/:productId/:afterId?", get.GetMessages)
	products.Get("posts/:productId/:field/:afterId?", get.GetPosts)
	products.Get("product/:product_name", get.GetProduct)
	products.Get("/latest_posts/:field", get.LatestProductPosts)
	products.Get("/dayStats/:productId", get.ProductDayStats)
	products.Get("board/:productId", get.GetFromPostsBoard)

	products.Put("board/:productId/:field/add/:postId", patch.AddPostToBoard)
	products.Put("board/:productId/:field/remove/:postId", patch.RemovePostFromBoard)
	products.Put("like_post/:postId", patch.LikePost)
	products.Put("dislike_post/:postId", patch.DislikePost)

	products.Post("/create", post.CreateProduct)
	products.Post("verify/:name", post.VerifyProduct)
	products.Post("/follow/:product_name", post.FollowProduct)
	products.Post("/unfollow/:product_name", post.UnfollowProduct)
	products.Post("create/comment/", post.CreateComment)
	products.Post("like/", post.LikeProduct)
	products.Post("dislike/", post.DislikeProduct)

	products.Patch("/update/:name", patch.UpdateProduct)
	products.Post("create/post/:name/:field", post.CreatePost)
	products.Patch("update/post/:name/:field/:post_id", patch.UpdatePost)

	products.Delete("delete/post/:productId/:field/:post_id", delete.DeletePost)
	products.Delete("delete/product/:id", delete.DeleteProduct)
	products.Patch("update/role/:product_id/:user_id", patch.UpdateUserRole)
}
