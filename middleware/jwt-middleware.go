package middleware

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Failed to get JWT SECRET")
	}
	return nil
}
