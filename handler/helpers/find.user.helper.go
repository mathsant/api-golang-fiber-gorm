package helpers

import (
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/model/entity"
)

func FindOneUser(user entity.User, idUser string) bool {
	err := database.DB.Debug().First(&user, idUser).Error

	if err != nil {
		return false
	}

	return true
}
