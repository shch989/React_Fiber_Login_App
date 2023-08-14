package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shch989/React_Fiber_Login_App/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	app.Post("/api/register", controllers.Register)
}
