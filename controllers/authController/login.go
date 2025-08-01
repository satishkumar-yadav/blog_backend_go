package authController

import (

	//"blogApp/database"
	"blogApp/controllers/handler"
	mid1 "blogApp/middleware/cookie"
	mid2 "blogApp/middleware/header"
	"blogApp/models"
	"strings"

	//"blogApp/utils"

	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

func Login2(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB) // c.Locals("db").(*YourDBType)  // db - key storing database

	//declaring variable/datatype to store parsed data come from http req
	//var data map[string]string // o
	var body struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}

	// parsing and storing to above variable
	//if err := c.BodyParser(&data); err != nil { // o
	if err := c.BodyParser(&body); err != nil {
		fmt.Println("Unable to parse body, Error : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please Enter Credential in Correct  Format"})
	}

	// creating model instance to store record return by db First operation
	var user models.User

	// check if userid exists in db
	// db.Where("user_id=?", data["user_id"]).First(&user) // o
	// if user.UserID == "" {
	// 	return c.Status(404).JSON(fiber.Map{"message": "UserID doesn't exit, kindly create an account"})
	// }
	if err := db.Where("user_id = ?", body.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid UserID !!"}) // "message": "Invalid Credential !!"
	}

	// check if password is correct for above userid in db
	//if err := user.ComparePassword(data["password"]); err != nil { // o
	if err := user.ComparePassword(body.Password); err != nil { // o
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Incorrect Password !!"}) // "message": "Invalid Credential !!"
	}

	// token based no cookie
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(mid2.SecretKey)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenStr, "message": "you have successfully login", "user": user})
}

func Login(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body with Error : ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please Provide Credential in Correct Format !"})
	}

	var user models.User
	//database.DB.Where("email=?", data["email"]).First(&user)

	if handler.ValidateEmail(strings.TrimSpace(data["user_id"])) {

		db.Where("email=?", data["user_id"]).First(&user)
		if user.UserID == "" {
			//c.Status(404)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Email Address doesn't exit, kindly create an account"})
		}

	} else {

		if err := db.Where("user_id = ?", data["user_id"]).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid UserID or doesn't exists !!, kindly create an account"})
		}

	}

	if err := user.ComparePassword(data["password"]); err != nil {
		//c.Status(400)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Incorrect Password"})
	}

	// token, err := mid1.GenerateJwt(strconv.Itoa(int(user.UserID)))  // o
	token, err := mid1.GenerateJwt(user.UserID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   false, // false for local, true for HTTPS
		SameSite: "Lax", // Optional but helpful, or "None" with Secure: true for cross-origin
		Path:     "/",   // ensure path is root
		//Domain:   "localhost",  // ✅ add this, hardcoding the backend cookie domain temporarily: If still blank
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Congrats! You have Successfully Logged-in",
		"user":    user,
		// "user-id":   user.UserID,
		// "user-name": user.FirstName,
		// "expires": time.Now().Add(time.Hour * 24).Unix(), // 👈 Send expiry  //////
		"expires": time.Now().Add(time.Hour * 24), // 👈 Send expiry  //////
	})

}
