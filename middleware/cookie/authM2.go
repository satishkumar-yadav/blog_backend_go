package cookie

import (
	"blogApp/config"
	"fmt"
	"time"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var SecretKey2 = []byte(config.SECRET_KEY)

// M-II
func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	issuer, err := Parsejwt(cookie)
	if err != nil {
		fmt.Println("error", err.Error())
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthenticated ! Token Expired or not found, Please Login First"}) // expose token expiration
	}

	c.Locals("userId", issuer) //
	return c.Next()
}

type Claims struct {
	jwt.StandardClaims
}

func Parsejwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey2), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

func GenerateJwt(issuer string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(SecretKey2))

}

// using custom claim
// func GenerateJWT(user models.User) (string, error) {

// 	claims := jwt.MapClaims{
// 		"email":  user.Email,
// 		"role":   user.Role,
// 		"course": user.Course,
// 		"exp":    time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(SecretKey2)
// }
