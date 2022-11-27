package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/model/entity"
	"mathsant/web-service-fiber/model/request"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})
}
