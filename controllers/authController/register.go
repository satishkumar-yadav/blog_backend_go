package authController

import (

	//"blogApp/database"
	"blogApp/controllers/handler"
	"blogApp/models"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	//var data map[string]interface{}  //o
	var body struct {
		UserID    string `json:"user_id"  `
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email" `
		Password  string `json:"password"`
		Phone     string `json:"phone"`
	}

	// if err := c.BodyParser(&data); err != nil { // o
	if err := c.BodyParser(&body); err != nil {
		fmt.Println("Unable to parse body with Error : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please Provide Information in Correct Format !"})
	}

	if body.UserID == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username and Password required"})
	}

	//o  status(400)
	// if !(handler.ValidateEmail(strings.TrimSpace(data["email"].(string)))) {
	if !(handler.ValidateEmail(strings.TrimSpace(body.Email))) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Email Address"})
	}

	var userData models.User // o
	//Check if email already exist in database
	// db.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	// Query for either email or user ID match
	// Prepare trimmed input
	email := strings.TrimSpace(body.Email)
	userID := strings.TrimSpace(body.UserID)

	db.Where("email = ? OR user_id = ?", email, userID).First(&userData)

	if userData.Email == email {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already exist",
		})
	} else if userData.UserID == userID && userID != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User ID already exist",
		})
	}

	//Check if password is less than 6 characters //o
	// if len(data["password"].(string)) <= 6 {
	if len(body.Password) <= 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Password must be greater than 6 character"})
	}

	// data["userid"] - is of types interface but UserID is of string so typecasted - data["userid"].(string)
	// user := models.User{
	// 	UserID:    data["user_id"].(string),
	// 	FirstName: data["first_name"].(string),
	// 	LastName:  data["last_name"].(string),
	// 	Phone:     data["phone"].(string),
	// 	Email:     strings.TrimSpace(data["email"].(string)),
	// }
	// user.SetPassword(data["password"].(string))

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		UserID:    body.UserID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash), // if used function then don't update here
		Phone:     body.Phone,
	}
	// user.SetPassword(body.Password) // o

	if err := db.Create(&user).Error; err != nil {
		log.Println("Data Creation Error : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Registration Failed!"})
	}

	// receiver := user.Email
	//   name := user.FirstName
	// error := utils.SendApprovalEmail(receiver, name) // smallLine of Code,working - email.go
	// if error != nil {
	// 	fmt.Println("Error sending email", error)
	//return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "You are Registered Successfully, but failed to send mail !"})
	// }

	//utils.SendGmail(receiver)  // giving error

	//fmt.Println("Registration Email Successfully sent")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "You are Registered Successfully"})
}
