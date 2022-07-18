package main

import (
    "log"
    
    "github.com/gofiber/fiber/v2"
)
func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("\nGoFiber Golang webframework\nI'm Running on Podman Port 8080\n\n")
    })

    log.Fatal(app.Listen(":8080"))
}
