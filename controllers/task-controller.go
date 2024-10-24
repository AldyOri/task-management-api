package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not create task",
		})
	}

	return c.JSON(http.StatusCreated, Response{Message: "task created", Data: task})
}

func GetTasks(c echo.Context) error {
	var tasks []models.Task

	if err := config.DB.Order("id").Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve tasks",
		})
	}

	return c.JSON(http.StatusOK, Response{Message: "task retrived", Data: tasks})
}

func GetTaskById(c echo.Context) error {
	var task models.Task
	id := c.Param("id")

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "task not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: "success",
		Data:    task,
	})
}

func UpdateTaskById(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	var updatedTask models.Task

	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "task not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve task",
		})
	}

	if err := config.DB.Model(&task).Updates(updatedTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Could not update task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: "task updated",
		Data:    task,
	})
}

func DeleteTaskById(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Task not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Could not retrieve task",
		})
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Could not delete task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Task deleted successfully",
		Data: task,
	})
}
