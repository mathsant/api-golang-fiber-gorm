package routes

import (
	"mathsant/web-service-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandleGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
}
