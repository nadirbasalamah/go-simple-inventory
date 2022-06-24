package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-simple-inventory/database"
	"github.com/nadirbasalamah/go-simple-inventory/routes"
)

const DEFAULT_PORT = "3000"

func main() {
	var app *fiber.App = fiber.New()

	database.InitDatabase()

	routes.SetupRoutes(app)

	var PORT string = os.Getenv("PORT")

	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))
}
