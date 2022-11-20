package handler

import (
	"log"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/handler/helpers"
	"mathsant/web-service-fiber/model/entity"
	"mathsant/web-service-fiber/model/request"

	"github.com/go-playground/validator/v10"
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

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()

	errValidator := validate.Struct(user)
	if errValidator != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidator.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newUser,
	})
}

func UserHandleGetById(ctx *fiber.Ctx) error {
	idUser := ctx.Params("id")

	var user entity.User

	err := database.DB.Debug().First(&user, idUser).Error

	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	userConverted := helpers.MakeUserResponse(&user)

	return ctx.JSON(userConverted)
}
