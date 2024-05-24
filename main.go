package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 80801
	log.Fatal(app.Listen(":8080"))
}
