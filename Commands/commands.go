package commands

import constants "example/main/Constants"


func CheckCommand(commands string, userstate map[int]string, IDmsg int) (string, bool) {
	var msg string
	var openButton bool
	switch commands {
		case "help":
			msg = "I understand /sayhi and /status."
		case "start":
			msg = constants.HelloMas
			openButton = true
		case "status":
			msg = "I'm ok."
		case "stop":
			msg = "Отключаем кнопки"
			openButton =false
		case "buy":
			userstate[IDmsg] = "awaiting_url"
			msg =  "Пришли мне ссылку на товар."
		case "delete":
			userstate[IDmsg] = "awaiting_url_for_delete"
			msg =  "Пришлите мне ссылку, чтобы удалить ее из моей забы данных"
		case "checkall":
			msg =  "Сейчас пришлю все то что надо купить))))"  
			userstate[IDmsg] = "return_all_url"
		default:
			msg = "что это за команда?"
	}
	return msg, openButton
}


