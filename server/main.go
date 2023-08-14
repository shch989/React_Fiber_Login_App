package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/React_Fiber_Login_App/database"
	"github.com/shch989/React_Fiber_Login_App/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
