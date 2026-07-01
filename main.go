package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lameesbeik/auth-acl-service/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterHealthRoutes(app)

	if err := app.Listen(":8110"); err != nil {
		log.Fatal(err)
	}
}
