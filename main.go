package main

import (
	commands "example/main/Commands"
	butoms "example/main/Buttons"
	openai "example/main/OpenAI"
	"log"
	"sync"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)


func main() {
	var buttons bool
	bot, err := tgbotapi.NewBotAPI("ТОКЕН БОТА!!!!")
	if err != nil {
		log.Panic(err)
	}

	wg := new(sync.WaitGroup)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 100

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
				msg.Text, buttons =  commands.CheckCommand(update.Message.Command())
				msg.ReplyMarkup = butoms.OpenOrCloseButton(buttons)
			}()
			wg.Wait()

        }else{
			log.Printf("[%s]/n %s/n", update.Message.From.UserName, update.Message.Text)
			response := openai.MessagefromGPT(update.Message.Text)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, response)
		}


        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
		
	}
}

