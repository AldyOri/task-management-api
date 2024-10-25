package routes

import (
	"todo-app/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	apiGroup.POST("/tasks", controllers.CreateTask)
	apiGroup.GET("/tasks", controllers.GetTasks)
	apiGroup.GET("/tasks/:id", controllers.GetTaskById)
	apiGroup.PATCH("/tasks/:id", controllers.UpdateTaskById)
	apiGroup.DELETE("/tasks/:id", controllers.DeleteTaskById)

	authGroup := apiGroup.Group("/auth")
	
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/me", controllers.GetMe)
}
