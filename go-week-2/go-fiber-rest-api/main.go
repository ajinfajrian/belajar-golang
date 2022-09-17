package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	database.DatabaseInit()
	app := fiber.New()
	// Initial Route
	route.RouteInit(app)
	app.Listen(":8080")
}
