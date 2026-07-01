package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lameesbeik/auth-acl-service/controllers"
)

func RegisterHealthRoutes(app *fiber.App) {
	app.Get("/", controllers.HealthCheck)
}
