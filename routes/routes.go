package routes

import (
	"todo-app/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/tasks", controllers.CreateTask)
	e.GET("/tasks", controllers.GetTasks)
	e.GET("/tasks/:id", controllers.GetTaskById)
	e.PATCH("/tasks/:id", controllers.UpdateTaskById)
	e.DELETE("/tasks/:id", controllers.DeleteTaskById)

	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.POST("/logout", controllers.LogOut)
}
