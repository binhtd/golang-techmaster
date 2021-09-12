package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	demoFiber1()
}

func demoFiber1() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
	app.Listen(3000)
}
