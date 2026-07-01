package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/JMustang/aurora_bank/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=aurora password=aurora123 dbname=aurora_bank port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar ao banco:", err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	fmt.Println("✅ Conexão com PostgreSQL realizada com sucesso!")
}
