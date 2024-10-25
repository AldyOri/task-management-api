package routes

import (
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/me", controllers.GetMe, middleware.JWTMiddleware())

	taskGroup := apiGroup.Group("/tasks", middleware.JWTMiddleware())

	taskGroup.POST("", controllers.CreateTask)
	taskGroup.GET("", controllers.GetTasks)
	taskGroup.GET("/:id", controllers.GetTaskById)
	taskGroup.PATCH("/:id", controllers.UpdateTaskById)
	taskGroup.DELETE("/:id", controllers.DeleteTaskById)
}
