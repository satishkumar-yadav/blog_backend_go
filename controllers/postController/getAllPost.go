package postController

import (
	"blogApp/models"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllPosts(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB) //
	/* Retrieves the GORM database connection from the request’s context local storage.
	   c.Locals("db") accesses data stored earlier in the middleware chain.
	   The .(*gorm.DB) part asserts the type—telling Go this is a pointer to gorm.DB.
	    This gives us access to the database for querying. */

	page, _ := strconv.Atoi(c.Query("page", "1")) //
	// c.Query("page", "1"): reads the page query parameter from the URL, or "1" if none provided.
	limit, _ := strconv.Atoi(c.Query("limit", "10")) //
	//limit := 5
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit //
	/*	Calculates the database query offset for pagination.
		offset is how many records to skip before starting to return results.
		Example: page=1 → offset=0; page=3, limit=10 → offset=20.
		Used to fetch the correct "slice" of data.
	*/

	var blogs []models.Blog
	/*	Declares a slice variable blogs to hold multiple Blog model instances.
		models.Blog is your application’s blog post struct/model representing the database table rows.
	*/

	// tx := db.Preload("User").Offset(offset).Limit(limit).Find(&blogs)  //o
	// tx := db.Offset(offset).Limit(limit).Preload("User").Find(&blogs)  // updated
	// tx := db.Offset(offset).Limit(limit).Order("created_at desc").Find(&blogs)
	tx := db.Preload("User").Offset(offset).Limit(limit).Order("created_at desc").Find(&blogs)
	if tx.Error != nil {
		fmt.Println("Error loading blog data : ", tx.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Failed to Load Blogs!"})
	}

	for i := range blogs {
		blogs[i].Image = c.BaseURL() + "/api/uploads/" + blogs[i].Image
	}
	/*	Iterates over all retrieved blogs.
		Modifies the Image field of each blog to be a full URL.
		c.BaseURL() returns the base URL of the server (e.g., http://localhost:3000).
		Concatenates /uploads/ plus the image filename stored in blogs[i].Image.
		This allows the client to receive absolute URLs for images, ready to load/display.
	*/

	var total int64 //
	db.Model(&models.Blog{}).Count(&total)
	/*	Declares total to hold total number of blog posts in the DB (ignores pagination).
		db.Model(&models.Blog{}).Count(&total) performs a SQL count(*) query on the blogs table.
		This is used for pagination metadata to let the frontend know the total count of blogs.
	*/

	last_page := math.Ceil(float64(int(total) / limit))
	fmt.Println(last_page)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blogs": blogs,
		"pagination": fiber.Map{
			"page": page, "limit": limit, "total": total, "last_page": last_page,
		},
	})
}
