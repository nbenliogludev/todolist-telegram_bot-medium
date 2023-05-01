package repositories

import (
	"telegram-todolist/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetTask(update tgbotapi.Update) error {
	task := models.Task{
		ChatId: update.Message.Chat.ID,
		Task:   update.Message.Text,
	}

	if result := DB.Create(&task); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTask(taskId string) error {
	if result := DB.Where("id = ?", taskId).Delete(&models.Task{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllTasks(chatId int64) ([]models.Task, error) {
	var tasks []models.Task
	if result := DB.Where("chat_id = ?", chatId).Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}
