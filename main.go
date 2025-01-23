package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        log.Println(c.BaseURL())
        log.Println(c.Get("Authorization"))
        return c.SendString("Hello World")
    })

    app.Listen(":8000")
}
