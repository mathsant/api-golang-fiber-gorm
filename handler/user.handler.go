package handler

import "github.com/gofiber/fiber/v2"

func UserHandlerRead(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(fiber.Map{
		"user": "Matheus",
	})
}
