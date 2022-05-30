package main

import (
	"GoSquish/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World")
	})
	routes.All(app)

	app.Listen("8080")
}
