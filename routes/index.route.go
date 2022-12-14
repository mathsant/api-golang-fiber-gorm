package routes

import (
	"mathsant/web-service-fiber/config"
	"mathsant/web-service-fiber/handler"
	"mathsant/web-service-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Post("/login", handler.LoginHandler)
	r.Get("/users", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandleGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Delete("/user/:id", handler.UserHandlerDelete)
	r.Get("/transactions", handler.TransactionHandlerGetAll)
	r.Post("/transactions", handler.TransactionHandlerCreate)
}
