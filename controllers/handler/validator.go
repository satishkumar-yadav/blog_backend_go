package handler

import (
	"regexp"
	//"github.com/dgrijalva/jwt-go"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

/*

func ValidatePhone(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func ValidatePassword(body interface{}) bool {
	//Check if password is less than 6 characters //o
	// if len(data["password"].(string)) <= 6 {
	if len(body.Password) <= 6 {
		return c.Status(fiber.StatusLengthRequired).JSON(fiber.Map{"message": "Password must be greater than 6 character"})
	}
}

func ValidateInput(body interface{}) bool {
	if body.UserID == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username and Password required"})
	}
}

func CheckInputExistence(body interface{}) bool {
	 db := c.Locals("db").(*gorm.DB)

     var userData models.User // o
	//Check if email already exist in database
	// db.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	db.Where("email=?", strings.TrimSpace(body.Email)).First(&userData)
	if userData.Email != "" {
		return c.Status(fiber.StatusIMUsed).JSON(fiber.Map{"message": "Email already exist"})
	}

	db.Where("user_id=?", strings.TrimSpace(body.UserID)).First(&userData)
	if userData.UserID != "" {
		return c.Status(fiber.StatusIMUsed).JSON(fiber.Map{"message": "User ID already exist"})
	}

}




*/
