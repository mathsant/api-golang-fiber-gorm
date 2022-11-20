package main

import (
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/database/migration"
	"mathsant/web-service-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":3000")
}
