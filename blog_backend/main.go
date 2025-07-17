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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // adjust based on frontend port
		AllowCredentials: true,
	}))

	routes.Setup(app)

	// start server
	// app.Listen(":"+port)
	log.Fatal(app.Listen(":" + port))

}
