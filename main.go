package main

import (
	"log"
	"todo-app/config"
	"todo-app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// @host localhost:8000
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description In value field type "Bearer" followed by a space and the JWT token. Otherwise you won't get authtorized, example : (Bearer <your-token>)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} \n",
	}))

	config.Connect()
	config.Migrate()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
