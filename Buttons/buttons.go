package buttons
import (
	"github.com/go-telegram-bot-api/telegram-bot-api"

)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("пососи писюн"),
        tgbotapi.NewKeyboardButton("писька"),
        tgbotapi.NewKeyboardButton("лох"),
    ),
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("бот"),
        tgbotapi.NewKeyboardButton("ты"),
        tgbotapi.NewKeyboardButton("/start"),
    ),
)
func OpenOrCloseButton(keyboard bool) interface{} {
	if keyboard {
	 	return numericKeyboard
	}else{
		return  tgbotapi.NewRemoveKeyboard(true)
	}
}