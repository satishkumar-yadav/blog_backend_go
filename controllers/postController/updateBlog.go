package postController

import (
	"blogApp/models"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	//"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateBlog(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("userId").(string)
	//userID := "abc3"
	id, _ := strconv.Atoi(c.Params("id"))

	var blog models.Blog
	if err := db.Preload("User").First(&blog, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Blog post not found"})
	}
	/*	- Declares a variable blog of type models.Blog to hold the fetched blog post.
		- db.First(&blog, id) tells GORM to fetch the first Blog record matching the primary key id.
		- .Error accesses any error that occurred in this query.
		- If no blog is found, returns an HTTP 404 Not Found response with a JSON error message.
		    - c.Status(fiber.StatusNotFound) sets the HTTP status code.
		    - .JSON() serializes the given map as JSON and sends it as the HTTP response.
	*/

	if blog.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "No Permission to Update, not your post"})
	}

	title := c.FormValue("title")
	desc := c.FormValue("description")
	file, _ := c.FormFile("image")
	/*   Reads form data from the HTTP request body:
	- c.FormValue("title") reads the form field "title" as a string.
	- c.FormValue("description") reads the form field "description".
	- c.FormFile("image") retrieves an uploaded file with the field name "image". Returns a *multipart.FileHeader or nil if no file uploaded.
	*/

	if title != "" {
		blog.Title = title
	}
	if desc != "" {
		blog.Description = desc
	}
	/*   Updates the blog's Title and Description only if new, non-empty values are provided in the form.
	This allows partial updates—fields left blank in the form won't overwrite existing values.
	*/

	fname := strings.ReplaceAll(file.Filename, " ", "")
	if file != nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fname)
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err == nil {
			blog.Image = filename
		}
	}
	/* If a file was uploaded (file != nil), process it:
	    - Constructs a unique filename by prepending the current Unix nanosecond timestamp to the original filename (reduces collisions).
	    - Uses c.SaveFile(file, path) to save the uploaded file to the local filesystem under the ./uploads/ directory.
	    - On successful save (err == nil), updates the blog's Image field with the new filename.
	This handles updating or replacing an existing blog image.
	*/

	db.Save(&blog)
	// result := db.Model(&blog).Updates(blog).Error
	// if result != nil {
	// 	return c.Status(fiber.StatusNotModified).JSON(fiber.Map{"message": "Blog update failed", "error": result})
	// }
	/*  Saves the updated blog record (including any changed fields) back to the database.
	    GORM's Save updates the existing record based on the primary key.
	*/

	// Return full image URL
	blog.Image = c.BaseURL() + "/api/uploads/" + blog.Image
	/*   Modifies blog.Image to be a full URL by prefixing the server's base URL and the /uploads/ folder path.
	c.BaseURL() returns the server's base URL (e.g., http://localhost:3000).
	This makes it convenient for clients to load/display the image from a full URL.
	*/

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "post updated successfully",
		"blogs":   blog,
	})
}

// o
func UpdatePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("userId").(string)
	//userID := "abc"

	// blog := models.Blog{
	// 	//Id: uint(id),
	// }

	var blog models.Blog
	if err := db.First(&blog, id).Error; err != nil {
		fmt.Println("Blog not Found Error : ", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Blog post not found"})
	}
	fmt.Println(c.JSON(blog))

	if blog.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "no permission to update, Not Your post"})
	}

	// Save the previous image filename
	prevImage := blog.Image

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Perform the update
	result := db.Model(&blog).Updates(blog).Error
	if result != nil {
		fmt.Println("Update Failed Error : ", result)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Blog update failed"})
	}

	// If image was changed and previous image existed, remove the old file
	if blog.Image != "" && prevImage != "" && blog.Image != prevImage {
		imagePath := filepath.Join(".", "uploads", prevImage)
		if err := os.Remove(imagePath); err != nil {
			fmt.Println("Old image delete error:", err)
			// Not critical, so don't fail the response
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "post updated successfully",
		"blogs":   blog,
	})

}
