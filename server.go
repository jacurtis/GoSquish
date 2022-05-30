package main

import (
	"GoSquish/db"
	"GoSquish/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db.Init()

	routes.All(app)

	app.Listen("8080")
}
