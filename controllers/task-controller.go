package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"
	"todo-app/models/dto"
	"todo-app/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task for the authenticated user
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param task body dto.TaskRequest true "Task"
// @Success 201 {object} models.Response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func CreateTask(c echo.Context) error {
	userID := utils.GetUserID(c)
	var task models.Task
	var taskRequest dto.TaskRequest

	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	task.UserID = userID
	task.Title = taskRequest.Title
	task.Description = taskRequest.Description
	task.Completed = taskRequest.Completed

	if err := config.DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not create task",
		})
	}

	return c.JSON(http.StatusCreated, dto.Response{Message: "task created", Data: dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}})
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get all tasks for the authenticated user
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetTasks(c echo.Context) error {
	userID := utils.GetUserID(c)
	var tasks []models.Task

	if err := config.DB.Where("user_id = ?", userID).Order("id").Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not retrieve tasks",
		})
	}

	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{Message: "task retrived", Data: taskResponses})
}

// GetTaskById godoc
// @Summary Get a task by ID
// @Description Get a task by ID for the authenticated user
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [get]
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

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success",
		Data: dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}

// UpdateTaskById godoc
// @Summary Update a task by ID
// @Description Update a task by ID for the authenticated user
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param task body dto.TaskRequest true "Task"
// @Success 200 {object} models.Response
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [patch]
func UpdateTaskById(c echo.Context) error {
	userID := utils.GetUserID(c)
	id := c.Param("id")
	var task models.Task
	var updatedTask dto.TaskRequest

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

	return c.JSON(http.StatusOK, dto.Response{
		Message: "task updated successfully",
		Data: dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}

// DeleteTaskById godoc
// @Summary Delete a task by ID
// @Description Delete a task by ID for the authenticated user
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [delete]
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

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Task deleted successfully",
		Data: dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		},
	})
}
