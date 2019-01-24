package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

type Boss struct {
	Bot   *tgbotapi.BotAPI
	Vacab Vacabruary
}

type Vacabruary struct {
	Qustions []string
	Unswers  []string
}

func main() {
	boss, err := newBoss(os.Getenv("TOKEN"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	log.Println("I`m ready")
	//Получаем обновления от бота
	updates := boss.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/say_hello":
				//Send message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a none_max_bot.")
				boss.Bot.Send(msg)
			default:
				q := rand.Intn(len(boss.Vacab.Qustions))
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(boss.Vacab.Qustions[q], update.Message.Text))
				boss.Bot.Send(msg)
				go func() {
					q = rand.Intn(len(boss.Vacab.Unswers))
					time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, boss.Vacab.Unswers[q])
					boss.Bot.Send(msg)
				}()
			}
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, `это что???`)
		boss.Bot.Send(msg)
	}
}

func newBoss(token string) (Boss, error) {
	var boss Boss
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return boss, err
	}
	boss.Bot = bot
	boss.Vacab.Qustions = getQustions()
	boss.Vacab.Unswers = getUnswers()
	return boss, nil
}

func getQustions() []string {
	return []string{
		0: `т.е. твой результат это "%s"?!`,
		1: `"%s", серьезно?`,
		2: `Меня не интересует это твое "%s", где результат?`,
		3: `Что значит "%s"?!`,
		4: `Постоянно "%s"!`,
	}
}

func getUnswers() []string {
	return []string{
		0: `Иди работай!`,
		1: `Мы это уже обсуждали!`,
		2: `Не хочу ничего слышать!`,
		3: `Ты скоро доиграешься`,
		4: `Я жду!`,
		5: `Мое терпение не бесконечно!`,
	}
}
