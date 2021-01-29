package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func printHello(c *fiber.Ctx) error {
	fmt.Printf("%+v\n", c)
	return c.SendString("Hello, World 👋!")
}
func main() {
	app := fiber.New()

	app.Get("/", printHello)

	app.Listen(":3000")
}
