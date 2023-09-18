package main

import (
	commands "example/main/Commands"
	butoms "example/main/Buttons"
	//openai "example/main/OpenAI"
	BD "example/main/DataBase"
	"log"
	"sync"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	constants "example/main/Constants"
	img "example/main/Img"

)

type UserState struct {
    UserID int
    State  string
}



func main() {
	var userStates = make(map[int]string) 
	var buttons bool
	bot, err := tgbotapi.NewBotAPI("5337023432:AAGWQ7HDN8mLTYGkeg8WbEdGefCnRuylHSw")
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
				msg.Text, buttons =  commands.CheckCommand(update.Message.Command(), userStates, update.Message.From.ID)
				msg.ReplyMarkup = butoms.OpenOrCloseButton(buttons)
			}()
			wg.Wait()

        }else if update.Message.Photo != nil{
			//работа с фотографиями
			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, tgbotapi.FileBytes(img.SendPhoto()))
			if _, err = bot.Send(photo); err != nil {
				log.Fatalln(err)
			}		
		} else{
			switch userStates[update.Message.From.ID] {
			case constants.StateAwaitingURL:
				url := update.Message.Text
				if !BD.IsAccessibleURL(url){
					msg.Text = "это не ссылка"
				}else{
					wg.Add(1)
					go func() {
						defer wg.Done()
							if BD.ConnectedForDB(url,constants.StateAwaitingURL) == 1{
							msg.Text = "ссылку принял..."
						}else if BD.ConnectedForDB(url,constants.StateAwaitingURL) == 404{
							msg.Text = "я не смог сохранить себе 😿"
						}else if BD.ConnectedForDB(url,constants.StateAwaitingURL) == 505 {
							msg.Text = "у меня уже есть эта ссылка"
						}
					}()
					wg.Wait()
				}
				userStates[update.Message.From.ID] = constants.StateNone
			case constants.StateAwaitingURLForDelete:
				url := update.Message.Text
				if !BD.IsAccessibleURL(url){
					msg.Text = "это не ссылка"
				}else{
					wg.Add(1)
					go func() {
						defer wg.Done()
							if BD.ConnectedForDB(url,constants.StateAwaitingURLForDelete) == 1{
							msg.Text = "Удалил)"
						}else if BD.ConnectedForDB(url,constants.StateAwaitingURLForDelete) == 404{
							msg.Text = "я не смог удалить, у меня проблемы😿"
						}else if BD.ConnectedForDB(url,constants.StateAwaitingURLForDelete) == 300 {
							msg.Text = "Ссылка не найдена"
						}
					}()
					wg.Wait()
				}
				userStates[update.Message.From.ID] = constants.StateNone
			case constants.StateReturnAllUrl:// тут должен возвращаться полный список всех ссылок
				
			default:
				log.Printf("[%s]/n %s/n", update.Message.From.UserName, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				// response := openai.MessagefromGPT(update.Message.Text)
				// msg = tgbotapi.NewMessage(update.Message.Chat.ID, response)
			}
		}

        if _, err := bot.Send(msg); err != nil {
            log.Print(err)
        }
		
	}
}

