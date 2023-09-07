package commands

import constants "example/main/Constants"



type CommandsInterface interface {
	CheckCommand(commands string) (string)
}

func CheckCommand(commands string) (string) {
	var msg string
	switch commands {
	case "help":
		msg = "I understand /sayhi and /status."
	case "start":
		msg = constants.HelloMas
	case "status":
		msg = "I'm ok."
	default:
		msg = "что это за команда?"
	}
	return msg
}
