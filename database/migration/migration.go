package migration

import (
	"fmt"
	"log"
	"mathsant/web-service-fiber/database"
	"mathsant/web-service-fiber/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
