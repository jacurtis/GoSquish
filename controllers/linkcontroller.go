package controllers

import "github.com/gofiber/fiber"

func GetLink(c *fiber.Ctx) {
	c.Send("google.com")
}

func PostLink(c *fiber.Ctx) {
	c.Send("Added new page")
}
