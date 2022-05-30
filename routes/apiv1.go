package routes

import (
	"GoSquish/controllers"
	"github.com/gofiber/fiber"
)

func ApiV1Routes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/link/:slug", controllers.GetLink)
	apiv1.Post("/link", controllers.PostLink)
	apiv1.Delete("/link/:slug", controllers.DeleteLink)
}
