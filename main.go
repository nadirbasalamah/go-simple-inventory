package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-simple-inventory/database"
	"github.com/nadirbasalamah/go-simple-inventory/routes"
)

func main() {
	var app *fiber.App = fiber.New()

	database.InitDatabase()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
