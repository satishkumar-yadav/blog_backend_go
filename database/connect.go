package database

import (
	"fmt"

	"blogApp/config"
	"blogApp/models"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func Connect(app *fiber.App) {
	dsn := config.DSN
	fmt.Println("DSN : ",dsn)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("DB Connection error: %v\n", err)
		panic("Could not connect to the database")
	} else {
		//log.Println("Database connected successfully")
		fmt.Println("Database connected successfully")
	}

	// DB = database
	database.AutoMigrate(
		&models.Blog{},
		&models.User{},
	)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", database) // db -key storing database
		return c.Next()
	})

}

/*
The Middleware Function: func(c *fiber.Ctx) error
This anonymous function is the middleware handler that takes a Fiber context c as its argument and returns an error.
fiber.Ctx encapsulates all the information about the HTTP request and provides methods to control the response.

c.Locals("db", db)
Locals is a method on the context used to store and retrieve local values during a request lifecycle.
Here, a value db (likely a database connection or instance) is stored in the request context under the key "db".
This means for this request, any subsequent handler or middleware can get access to the database by calling c.Locals("db").

return c.Next()
This tells Fiber to continue to the next middleware/handler in the chain.
If you don't call c.Next(), the chain stops here and the request won't proceed to other handlers or your route.

*/
