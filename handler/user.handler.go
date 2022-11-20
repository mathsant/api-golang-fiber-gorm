package handler

import (
	"log"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/model/entity"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Debug().Find(&users)
	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(users)
}
