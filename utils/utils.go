package utils

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	return uint(userID)
}

func GetTaskID(c echo.Context) (uint, error) {
	taskIDStr := c.Param("task_id")

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil{
		return 0, err
	}

	return uint(taskID), nil
}
