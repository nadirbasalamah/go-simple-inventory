package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-simple-inventory/handlers"
	"github.com/nadirbasalamah/go-simple-inventory/middlewares"
)

func SetupRoutes(app *fiber.App) {
	// public routes
	app.Post("/api/v1/signup", handlers.Signup)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/items", handlers.GetAllItems)
	app.Get("/api/v1/items/:id", handlers.GetItemByID)

	// private routes, authentication is required
	app.Post("/api/v1/items", middlewares.CreateMiddleware(), handlers.CreateItem)
	app.Put("/api/v1/items/:id", middlewares.CreateMiddleware(), handlers.UpdateItem)
	app.Delete("/api/v1/items/:id", middlewares.CreateMiddleware(), handlers.DeleteItem)

	// testing only
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		if id == "1234" {
			c.Cookie(&fiber.Cookie{
				Name:    "CookieForAndy",
				Value:   "Andy",
				Expires: time.Now().Add(24 * time.Hour),
			})
			return c.JSON(map[string]string{"id": "1234", "name": "Andy"})
		}

		return c.SendStatus(http.StatusNotFound)
	})
}
