package postController

import (
	"blogApp/models"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateBlog(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("userId").(string)

	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return err
	// }
	// files := form.File["image"]

	// Expect fields: title, description (as form fields), image file (as multipart file)
	title := c.FormValue("title")
	desc := c.FormValue("description")
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("error reading file : ", err)
	}

	// var filename string
	// if len(files) > 0 {
	// 	file := files[0]
	// 	filename = fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	// 	if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err != nil {
	// 		return err
	// 	}
	// }

	fname := strings.ReplaceAll(file.Filename, " ", "")
	//fmt.Println("Space Removed : ", fname)

	filename := ""
	if file == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "File not received"})
	} else {
		filename = fmt.Sprintf("%d_%s", time.Now().UnixNano(), fname)
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Could not Upload file on network"})
		}
	}
	// if file != nil {   }

	if title == "" || desc == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing required fields:  Title,  Description"})
	}

	// blog := models.Blog{Title: form.Value["title"][0], Description: form.Value["description"][0], Image: filename, UserID: userID}
	blog := models.Blog{Title: title, Description: desc, Image: filename, UserID: userID}
	//db.Preload("User").Create(&blog) // preload not working
	db.Create(&blog)

	blog.Image = c.BaseURL() + "/api/uploads/" + filename

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Congratulation!, Your Blog is Posted Successfully and is live now.",
	})
}

func CreatePost(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var blogpost models.Blog

	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}

	fmt.Println("parsed data: ", c.JSON(blogpost))

	if err := db.Create(&blogpost).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error creating blog post"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Congratulation!, Your post is live"})

}
