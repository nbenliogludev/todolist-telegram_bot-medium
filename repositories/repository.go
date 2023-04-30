package repositories

import (
	"telegram-todolist/database"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
