package main

import (
	"github.com/gofiber/fiber/v2"
	"logger_service/internal/api"
	"logger_service/internal/config"
	"logger_service/internal/utils"
)

func main() {

	// Initialize environment variables
	config.InitEnvConfigs()
	// Initialize database source name

	app := fiber.New(fiber.Config{
		//ReadTimeout:   time.Second * 15,
		//WriteTimeout:  time.Second * 15,
		Concurrency:  10,
		ServerHeader: "logger_service_v1",
		AppName:      "logger_service_v1",
	})

	// Register routes
	api.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(config.EnvConfigs.LocalServerPort)
	utils.CheckErr("Error while serving the api", err)

}
