package api

import (
	"go-e-commerce/config"

	"github.com/gofiber/fiber/v3"
)


func StartServer(config config.AppConfig){
	app := fiber.New()


	app.Listen(config.ServerPort)
}