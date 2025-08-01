package postController

import (
	"blogApp/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Optional helper: view one post (by owner only) - detail post
func MyPostDetail(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("userId").(string)
	//userID := "abc3"
	id, _ := strconv.Atoi(c.Params("id"))

	var blog models.Blog
	if err := db.Where("id=?", id).Preload("User").First(&blog).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	if blog.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Not Your Post"})
	}

	blog.Image = c.BaseURL() + "/api/uploads/" + blog.Image

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"blogs": blog})
}
