package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	ChatId int64     `gorm:"chat_id"`
	Task   string    `gorm:"task"`
}
