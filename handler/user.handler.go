package handler

import (
	"log"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/handler/helpers"
	"mathsant/web-service-fiber/model/entity"
	"mathsant/web-service-fiber/model/request"
	"mathsant/web-service-fiber/model/response"

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

	userFound := helpers.FindOneUser(user, idUser)

	if !userFound {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	userConverted := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(userConverted)
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	err := database.DB.First(&user, userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	if userRequest.Phone != "" {
		user.Phone = userRequest.Phone
	}
	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal save error for update",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	idUser := ctx.Params("id")

	var user entity.User

	userFound := helpers.FindOneUser(user, idUser)

	if !userFound {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	errDelete := database.DB.Debug().Delete(&user, idUser).Error

	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal save error for update",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User deleted",
	})
}
