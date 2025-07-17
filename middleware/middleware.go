package middleware

import (
	"blog/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	issuer, err := util.Parsejwt(cookie)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
			"error":   err.Error(), // expose token expiration
		})
	}
	c.Locals("userID", issuer) //
	return c.Next()

}
