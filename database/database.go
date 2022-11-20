package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseInit() {
	stringOfConnection := "host=localhost user=root password=root dbname=fibergo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringOfConnection))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	fmt.Println("Connecto to database")
}
