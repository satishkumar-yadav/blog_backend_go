package routes

import (
	auth "blogApp/controllers/authController"
	"blogApp/controllers/imageController"
	post "blogApp/controllers/postController"

	mid1 "blogApp/middleware/cookie"

	//mid2 "blogApp/middleware/header"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with Auto-Reload!")
	})
	// api endpoint example
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pong-pong!"})
	})

	// AUTH
	//app.Post("/forget/password", auth.ForgetPassword)
	//app.Post("/forget/userid", auth.ForgetUserId)
	app.Post("/auth/register", auth.Register) //w
	app.Post("/auth/login", auth.Login)       //w cookie based // can login either with userid or email with password
	//app.Post("/auth/login", auth.Login2)     //w token header based

	app.Get("/api/blogs", post.GetAllPosts)    //w
	app.Get("/api/blogs/:id", post.DetailPost) //w

	app.Static("/api/uploads", "./uploads") // w  //for serving uploaded images

	// BLOGS - all require Auth
	//app.Use(mid1.IsAuthenticate)
	//api := app.Group("/api")
	//api := app.Group("/api", mid2.AuthRequired) //w
	api := app.Group("/api", mid1.IsAuthenticate) //w

	api.Post("/logout", auth.Logout) //w
	//api.Post("/logout", auth.Logout2) //only message

	//api.Post("/create-blog", post.CreateBlog) //w
	api.Post("/create-blog", post.CreatePost) // w // when uploaded image url available

	//api.Put("/update-blog/:id", post.UpdateBlog) //w
	api.Put("/update-blog/:id", post.UpdatePost) //w  // when uploaded image url available

	api.Delete("/delete-blog/:id", post.DeleteBlog) //w

	api.Get("/myblogs", post.GetMyPosts)       //w
	api.Get("/myblogs/:id", post.MyPostDetail) //w

	api.Post("/upload-image", imageController.Upload) //w

	app.Get("/api/check-auth", func(c *fiber.Ctx) error { //w
		cookie := c.Cookies("jwt")
		return c.JSON(fiber.Map{"jwt": cookie})
	})

}
