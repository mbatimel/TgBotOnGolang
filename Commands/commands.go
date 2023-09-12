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
            
	default:
		msg = "что это за команда?"
	}
	return msg, openButton
}


/*Проверка ссылок может быть выполнена на разных уровнях сложности в зависимости от ваших требований:

1. **Базовая проверка с помощью регулярных выражений**:
Можно использовать регулярное выражение для проверки, является ли текст URL-адресом.

```go
import "regexp"

func isValidURL(url string) bool {
    re := regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
    return re.MatchString(url)
}
```

2. **Проверка доступности URL**:
Если вы хотите удостовериться, что URL не только выглядит как URL, но и является действующим веб-сайтом, можно отправить запрос по этому адресу.

```go
import (
    "net/http"
    "time"
)

func isAccessibleURL(url string) bool {
    timeout := time.Duration(5 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }

    resp, err := client.Get(url)
    if err != nil {
        return false
    }
    defer resp.Body.Close()

    return resp.StatusCode == 200
}
```

    */