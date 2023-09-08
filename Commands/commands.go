package commands

import constants "example/main/Constants"



type CommandsInterface interface {
	CheckCommand(commands string) (string)
}

func CheckCommand(commands string) (string, bool) {
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
	default:
		msg = "что это за команда?"
	}
	return msg, openButton
}
