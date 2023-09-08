package buttons
import (
	"github.com/go-telegram-bot-api/telegram-bot-api"

)
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
func OpenOrCloseButton(keyboard bool) interface{} {
	if keyboard {
	 	return numericKeyboard
	}else{
		return  tgbotapi.NewRemoveKeyboard(true)
	}
}