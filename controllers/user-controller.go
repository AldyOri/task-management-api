package controllers

import (
	"net/http"
	"os"
	"time"
	"todo-app/config"
	"todo-app/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user models.User

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Invalid credentials. Please check your email",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid credentials. Please check your password.",
		})
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to get JWT secret",
		})
	}
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to hash password",
		})
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "user registered successfully",
	})
}

func GetMe(c echo.Context) error {
	return nil
}