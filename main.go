package main

import (
	"log"

	"blog/config"
	"blog/database"
	"blog/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	port := config.PORT

	app := fiber.New()

	// cors added - allow frontend from Vite port
	app.Use(cors.New(cors.Config{ // allowed both local and netlify
		AllowOrigins:     "https://go-blogmanager.netlify.app, http://localhost:5173", // adjust based on frontend port
		AllowCredentials: true,                                                        // if using cookies or jwt in credentials
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.Setup(app)

	// start server
	// app.Listen(":"+port)
	log.Fatal(app.Listen(":" + port))

}
