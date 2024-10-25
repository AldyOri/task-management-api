package routes

import (
	"todo-app/controllers"
	_ "todo-app/docs"
	"todo-app/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRoutes(e *echo.Echo) {
    e.GET("/swagger/*", echoSwagger.WrapHandler)


	apiGroup := e.Group("/api")

	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/register", controllers.Register)
	authGroup.GET("/me", controllers.GetMe, middleware.JWTMiddleware())

	taskGroup := apiGroup.Group("/tasks", middleware.JWTMiddleware())

	taskGroup.POST("", controllers.CreateTask)
	taskGroup.GET("", controllers.GetTasks)
	taskGroup.GET("/:id", controllers.GetTaskById)
	taskGroup.PATCH("/:id", controllers.UpdateTaskById)
	taskGroup.DELETE("/:id", controllers.DeleteTaskById)
}
