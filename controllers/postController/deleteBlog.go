package postController

import (
	"blogApp/models"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	//"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteBlog(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("userId").(string)
	//userID := "abc"
	fmt.Println("user id : ", userID)

	id, _ := strconv.Atoi(c.Params("id"))

	//fmt.Println("id : ", id)

	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Blog Not found"})
	}
	//fmt.Println("Blog Data : ", blog)

	if blog.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "No Permission to delete this"})
	}

	// Store the image filename before deleting the blog
	imageFile := blog.Image // Make sure your Blog model has an Image field

	// Delete blog entry from DB
	deleteQuery := db.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Opps!, record Not found"})
	}

	// Only proceed to delete the image if blog deletion was successful
	if deleteQuery.Error == nil && imageFile != "" {
		// Construct full path to the image file
		imagePath := filepath.Join(".", "uploads", imageFile)
		// Attempt to remove the file
		if err := os.Remove(imagePath); err != nil {
			// Log the error but don't fail the response
			fmt.Println("Failed to delete image file:", err)
		}
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Blog Deleted"})
}
