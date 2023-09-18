package img

import (
	"io/ioutil"
"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendPhoto() tgbotapi.FileBytes {
	photoBytes, err := ioutil.ReadFile("Img/imgStorage/parrot.png")
if err != nil {
    panic(err)
}
photoFileBytes := tgbotapi.FileBytes{
    Name:  "picture",
    Bytes: photoBytes,
}
	return photoFileBytes
}