package services

import (
	"fmt"

	"telegram-todolist/keyboards"
	"telegram-todolist/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hi, here you can create todos for your todolist."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, write todo."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Todo successfully created"

	err := repositories.SetTask(update)
	if err != nil {
		text = "Couldnt set task"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	data, _ := repositories.GetAllTasks(update.Message.Chat.ID)
	var btns []tgbotapi.InlineKeyboardButton
	for i := 0; i < len(data); i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Task, "delete_task="+data[i].ID.String())
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	//keyboard.InlineKeyboard = rows

	text := "Please, select todo you want to delete"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
	text := "Task successfully deleted"

	err := repositories.DeleteTask(taskId)
	if err != nil {
		text = "Couldnt delete task"
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Tasks: \n"

	tasks, err := repositories.GetAllTasks(update.Message.Chat.ID)
	if err != nil {
		text = "Couldnt get tasks"
	}

	for i := 0; i < len(tasks); i++ {
		text += tasks[i].Task + " \n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
