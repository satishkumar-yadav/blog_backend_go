package routes

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)

	app.Post("/api/logout", controller.Logout)
	app.Post("/api/create-blog", controller.CreatePost)       // /api/post
	app.Get("/api/get-blog", controller.AllPost)              // /allpost
	app.Get("/api/get-blog/:id", controller.DetailPost)       // /allpost/:id
	app.Put("/api/update-blog/:id", controller.UpdatePost)    // /updatepost
	app.Get("/api/unique-blog", controller.UniquePost)        //  /uniquepost
	app.Delete("/api/delete-blog/:id", controller.DeletePost) // /deletepost
	app.Post("/api/upload-image", controller.Upload)          //
	app.Get("/api/check-auth", func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		return c.JSON(fiber.Map{"jwt": cookie})
	})
	app.Static("/api/uploads", "./uploads")

}
