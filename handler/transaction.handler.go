package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/model/entity"
)

func TransactionHandlerGetAll(ctx *fiber.Ctx) error {
	var transactions []entity.Transaction
	err := database.DB.Debug().Preload("User").Find(&transactions)
	if err != nil {
		log.Println(err)
	}

	return ctx.JSON(transactions)
}

func TransactionHandlerCreate(ctx *fiber.Ctx) error {
	transaction := new(entity.Transaction)

	if err := ctx.BodyParser(transaction); err != nil {
		return err
	}

	errCreateTransaction := database.DB.Create(&transaction)
	if errCreateTransaction != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    transaction,
	})

}
