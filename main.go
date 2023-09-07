package main

import (
	commands "example/main/Commands"
	"log"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

//кнопки
var numericKeyboard = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("1"),
        tgbotapi.NewKeyboardButton("2"),
        tgbotapi.NewKeyboardButton("3"),
    ),
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("4"),
        tgbotapi.NewKeyboardButton("5"),
        tgbotapi.NewKeyboardButton("6"),
    ),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5337023432:AAGWQ7HDN8mLTYGkeg8WbEdGefCnRuylHSw")
	if err != nil {
		log.Panic(err)
	}
	wg := new(sync.WaitGroup)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Print(err)
	}

	for update := range updates {
		if update.Message == nil { 
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,update.Message.Text)

		if update.Message.IsCommand() { 
			wg.Add(1)
			
			go func() {
				defer wg.Done()
				//тут у меня бот отвечает на команды
				msg.Text =  commands.CheckCommand(update.Message.Command())
			}()
			wg.Wait()

        }else{
			//тут будет ответ на простые вопросы пользователя
			log.Printf("[%s]/n %s/n", update.Message.From.UserName, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
		}

		//открытие кнопок
		switch update.Message.Text {
        case "open":
            msg.ReplyMarkup = numericKeyboard
        case "close":
            msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
        }

        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
		
	}
}

