package main

import (
	"log"
	"os"
	"reflect"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	bot, err := newBot(os.Getenv("TOKEN"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	log.Println("I`m ready")
	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/say_hello":
				//Send message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a none_max_bot.")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Что значит "`+update.Message.Text+`"?!`)
				bot.Send(msg)
			}
		}
	}
}

func newBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return bot, err
	}
	return bot, nil
}
