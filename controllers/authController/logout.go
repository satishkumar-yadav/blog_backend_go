package authController

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Since there is no session or server-side token tracking, you usually do NOT need a server-side logout endpoint
// To "logout", your frontend (browser code) should simply remove the token from localStorage: localStorage.removeItem("token");
func Logout2(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

func Logout(c *fiber.Ctx) error {
	// c.ClearCookie("jwt")  // return
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",                             // clear cookie value
		Expires:  time.Now().Add(-1 * time.Hour), // expire immediately
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out Successfully"})
}
