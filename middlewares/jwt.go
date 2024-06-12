package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lmhuong711/go-go-be/utils"
)

var ROLES = []string{"user", "admin"}

func AuthToken(c *fiber.Ctx) error {
	jwt_secret := "jwt_secret"
	if os.Getenv("jwt_secret") != "" {
		jwt_secret = os.Getenv("jwt_secret")
	}

	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid token",
			"data":    nil,
			"count":   0,
		})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
			"count":   0,
		})
	}

	user := map[string]interface{}{
		"name": "",
		"role": "",
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	utils.InfoLog.Println("claims", claims)
	if !ok || !tokenByte.Valid ||
		claims["role"] == nil ||
		!slices.Contains(ROLES, claims["role"].(string)) {
		goto unauth
	}
	user["name"] = claims["name"].(string)
	user["role"] = claims["role"].(string)
	c.Locals("user", user)

	if user["role"] == "admin" {
		return c.Next()
	}
	if user["role"] == "user" && c.Method() == http.MethodGet {
		return c.Next()
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"success": false,
		"message": "Invalid role",
		"data":    nil,
		"count":   0,
	})

unauth:
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "Invalid token",
		"data":    nil,
		"count":   0,
	})

}
