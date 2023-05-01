package database

import (
	"fmt"
	"log"
	"telegram-todolist/config"
	"telegram-todolist/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	//Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_PORT"), "postgres")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec("DROP DATABASE IF EXISTS todolist;").Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("CREATE DATABASE todolist").Error; err != nil {
		panic(err)
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DATABASE_NAME"), config.Config("POSTGRES_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate tables
	DB.AutoMigrate(&models.Task{})

	return DB
}
