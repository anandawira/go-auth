package main

import (
	"github.com/anandawira/go-auth/database"
	"github.com/anandawira/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
