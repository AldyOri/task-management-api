package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"
	"todo-app/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTask(c echo.Context) error {
	userID := utils.GetUserID(c)
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	task.UserID = userID

	if err := config.DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not create task",
		})
	}

	return c.JSON(http.StatusCreated, models.Response{Message: "task created", Data: models.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}})
}

func GetTasks(c echo.Context) error {
	userID := utils.GetUserID(c)
	var tasks []models.Task

	if err := config.DB.Where("user_id = ?", userID).Order("id").Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve tasks",
		})
	}

	var taskResponses []models.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, models.Response{Message: "task retrived", Data: taskResponses})
}

func GetTaskById(c echo.Context) error {
	var task models.Task
	userID := utils.GetUserID(c)
	id := c.Param("id")

	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "task not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve task",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success",
		Data: models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}

func UpdateTaskById(c echo.Context) error {
	userID := utils.GetUserID(c)
	id := c.Param("id")
	var task models.Task
	var updatedTask models.Task

	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&task).Error; err != nil {
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

	return c.JSON(http.StatusOK, models.Response{
		Message: "task updated successfully",
		Data: models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}

func DeleteTaskById(c echo.Context) error {
	userID := utils.GetUserID(c)
	id := c.Param("id")
	var task models.Task

	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&task).Error; err != nil {
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

	return c.JSON(http.StatusOK, models.Response{
		Message: "Task deleted successfully",
		Data: models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}
