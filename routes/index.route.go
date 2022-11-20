package routes

import (
	"mathsant/web-service-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)
}
