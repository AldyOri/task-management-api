package middleware

import (
	"log"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Failed to get JWT SECRET")
	}
	return echojwt.JWT([]byte(jwtSecret))
}
