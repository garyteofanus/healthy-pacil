package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"healthy-pacil/database"
	"healthy-pacil/handlers"
	"log"
	"os"
)

func main() {
	// Setup Database
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())

	// Setup static files
	app.Static("/", "../client/build")

	api := app.Group("/api")

	api.Post("/register", handlers.Register)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen(":" + port))
}
