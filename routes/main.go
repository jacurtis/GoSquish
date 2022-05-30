package routes

import "github.com/gofiber/fiber"

func All(app *fiber.App) {
	app.Get("/hello", func(ctx *fiber.Ctx) {
		ctx.Send("Hello World")
	})
	ApiV1Routes(app)

}
