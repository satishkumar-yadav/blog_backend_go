package imageController

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrsuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	//fmt.Println("Form data : ", form)
	files := form.File["image"]
	//fmt.Println("Files data : ", files)
	fileName := ""

	if len(files) <= 0 {
		//fmt.Println("no files")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No Files Found, Re-Upload Again!"})
	}

	for _, file := range files {
		fname := strings.ReplaceAll(file.Filename, " ", "")
		fileName = randLetter(5) + "-" + fname
		//fmt.Println("File name : ", fileName)
		if er := c.SaveFile(file, "./uploads/"+fileName); er != nil {
			fmt.Println("File Saving Error : ", er)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error Uploading File on Network"})
		}
	}

	//fmt.Println("uploaded ")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "File Uploaded Successfully",
		"image_url": c.BaseURL() + "/api/uploads/" + fileName,
		"filename":  fileName,
	})

}
