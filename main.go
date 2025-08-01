package main

import (
	"blogApp/config"
	"blogApp/database"
	"blogApp/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create uploads directory for images
	os.MkdirAll("./uploads", os.ModePerm)
	port := config.PORT

	app := fiber.New()

	database.Connect(app)

	//app.Use(logger.New())
	// alternative for above code
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${method} ${path} - ${status} - ${latency}\n",
		TimeZone:   "Local",
		TimeFormat: "02-Jan-2006 15:04:05",
	})) //  tailor the log output format, time zone, and other behaviors.

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173 , https://go-blogmanager.netlify.app",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	// start server
	fmt.Println("Server started on - http://localhost:" + port)
	//log.Fatal(app.Listen(":8080"))
	log.Fatal(app.Listen(":" + port))
}

/*
	// cors added - allow frontend from Vite port
	app.Use(cors.New(cors.Config{ // allowed both local and netlify
		AllowOrigins:     "https://go-blogmanager.netlify.app, http://localhost:5173", // adjust based on frontend port
		AllowCredentials: true,                                                        // if using cookies or jwt in credentials
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))
*/

/*

app := fiber.New()
	// fundamental starting point for creating a new Fiber web application instance in Go.
	// fiber.New() - This function initializes and returns a new Fiber application. It sets up the internal router, middleware stack, configuration defaults, and everything required to start accepting and handling HTTP requests.
	// app := - This assigns the newly created Fiber app instance to the app variable, which you use to define routes, register middleware, and start the server.

	app.Use(logger.New())
	//app.Use(logger.New()) registers a middleware that logs request/response info for every HTTP request.
	//It’s automatic, lightweight, and helps you see what’s happening inside your Fiber app.

	// alternative for above code
	// 	app.Use(logger.New(logger.Config{
	//     Format:     "${time} - ${method} ${path} - ${status} - ${latency}\n",
	//     TimeZone:   "Local",
	//     TimeFormat: "02-Jan-2006 15:04:05",
	// }))  //  tailor the log output format, time zone, and other behaviors.

	// app.Use(...) - This function is a middleware registration method in Fiber. Middleware is a function that gets executed for every incoming HTTP request before reaching the final route handler (or next middleware).
    // app.Use(handler) registers the given handler function to be applied on all requests (or requests matching some criteria if a path or method is specified).


*/
