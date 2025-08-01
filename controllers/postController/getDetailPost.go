package postController

import (
	"blogApp/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// add pagination
func DetailPost(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id, _ := strconv.Atoi(c.Params("id"))

	var blogpost models.Blog
	if err := db.Where("id=?", id).Preload("User").First(&blogpost).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Blog Not found !"})
	}

	// add image link
	blogpost.Image = c.BaseURL() + "/api/uploads/" + blogpost.Image

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blogs": blogpost,
	})

}
