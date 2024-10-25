package controllers

import (
	"net/http"
	"os"
	"time"
	"todo-app/config"
	"todo-app/models"
	dto "todo-app/models/dto"
	"todo-app/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func Login(c echo.Context) error {
	var body dto.LoginRequest
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
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":     time.Now().Unix(),
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

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param register body models.User true "Register"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
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

// GetMe godoc
// @Summary Get current user
// @Description Get details of the authenticated user
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response
// @Failure 500 {object} map[string]string
// @Router /auth/me [get]
func GetMe(c echo.Context) error {
	userID := utils.GetUserID(c)

	var user models.User
	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve user",
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success",
		Data: dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
