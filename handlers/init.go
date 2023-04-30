package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// Loop through each update.
	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message.IsCommand() {
			Commands(bot, update)
		} else {
			Messages(bot, update)
		}
	}
}
