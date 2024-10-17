package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/api", func (c *fiber.Ctx) error {
        return c.SendString("Hello, Backend Server")
    })

    log.Fatal(app.Listen(":3000"))
}