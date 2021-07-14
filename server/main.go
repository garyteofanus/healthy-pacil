package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"healthy-pacil/database"
	"healthy-pacil/handler"
	"log"
	"os"
)

func main() {
	// Setup Database
	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup static files
	app.Static("/", "../client/build")

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/logout", handler.Logout)

	appointment := api.Group("/appointment")
	appointment.Get("/", handler.GetAllAppointments)
	appointment.Post("/:id/apply", handler.ApplyAppointment)
	appointment.Post("/:id/cancel", handler.CancelAppointment)
	appointment.Post("/create", handler.CreateAppointment)
	appointment.Put("/:id/update", handler.UpdateAppointment)
	appointment.Delete("/:id/delete", handler.DeleteAppointment)
	appointment.Get("/:id/patients", handler.GetPatientsFromAppointment)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	var port string
	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}

	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen(":" + port))
}
