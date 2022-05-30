package routes

import (
	"GoSquish/controllers"
	"github.com/gofiber/fiber"
)

func ApiV1Routes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/link", controllers.GetLink)
	apiv1.Post("/link", controllers.PostLink)
}
