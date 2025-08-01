package postController

import (
	"blogApp/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyPosts(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	// cookie := c.Cookies("jwt")  //o
	// userID, _ := utils.Parsejwt(cookie)

	userID := c.Locals("userId").(string)
	//fmt.Println(userID)
	//userID := "abc"

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var blogs []models.Blog
	// tx := db.Model(&blogs).Where("user_id=?", userID).Preload("User").Find(&blogs)  //o
	tx := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Order("created_at desc").Preload("User").Find(&blogs)
	if tx.Error != nil {
		fmt.Println("No Blog Error : ", tx.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No Blog found !, Kindly Create Blog First."})
	}

	//Patch image URL
	for i := range blogs {
		blogs[i].Image = c.BaseURL() + "/api/uploads/" + blogs[i].Image
	}

	var total int64
	db.Model(&models.Blog{}).Where("user_id = ?", userID).Count(&total)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blogs": blogs,
		"pagination": fiber.Map{
			"page": page, "limit": limit, "total": total,
		},
	})
}
