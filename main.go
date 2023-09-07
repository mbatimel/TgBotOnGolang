package main

import (
	commands "example/main/Commands"
	"log"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
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

        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
		
	}
}

