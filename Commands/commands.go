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

/*Чтобы заставить вашего бота в Telegram дожидаться ответа пользователя на Golang, вам нужно использовать состояния (state machine) или, другими словами, управление состоянием пользователя. Ниже приведен пример, как это можно реализовать:

1. **Хранение состояний**:
Для начала вам нужен механизм для хранения текущего состояния каждого пользователя. Это может быть база данных, in-memory storage или другие подходы.

```go
type UserState struct {
    UserID int
    State  string
}

var userStates = make(map[int]string) // простой пример in-memory storage
```

2. **Определение состояний**:
Вам нужно определить, какие состояния будут у вашего бота. Например:
```go
const (
    StateNone        = "none"
    StateAwaitingURL = "awaiting_url"
)
```

3. **Обработка команд и сообщений**:
При обработке каждого сообщения от пользователя вы должны проверять его текущее состояние и действовать соответственно.

```go
func handleUpdate(update tgbotapi.Update) {
    if update.Message.IsCommand() {
        switch update.Message.Command() {
        case "добавить":
            userStates[update.Message.From.ID] = StateAwaitingURL
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пришли мне ссылку на товар.")
            bot.Send(msg)
        }
    } else {
        switch userStates[update.Message.From.ID] {
        case StateAwaitingURL:
            url := update.Message.Text
            // здесь можете сделать проверку ссылки и т.д.
            // ...
            userStates[update.Message.From.ID] = StateNone
        default:
            // обработка обычных сообщений
        }
    }
}
```

Таким образом, когда пользователь отправляет команду `/добавить`, бот меняет состояние пользователя на `StateAwaitingURL` и ожидает ссылку. Когда ссылка приходит, бот обрабатывает ее в соответствующем состоянии, а затем возвращает пользователя в нормальное состояние.

Данный пример - это базовая идея, которую вы можете доработать и адаптировать под свои нужды.*/


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

3. **Проверка безопасности URL**:
Если вы хотите удостовериться, что URL безопасен (например, не ведет на фишинговый сайт), можно использовать внешние сервисы, такие как Google Safe Browsing.

4. **Проверка содержимого URL**:
В зависимости от вашего приложения, вы можете захотеть проверить, что URL ведет на определенный тип содержимого. Например, если вы ожидаете, что это будет изображение, вы можете загрузить содержимое и проверить его MIME-тип.

5. **Дополнительные проверки**:
Вы можете добавить дополнительные проверки, такие как проверка домена на наличие в черном списке, проверка на использование URL-shorteners и другие.

Какую именно проверку выбирать, зависит от вашей задачи и того, насколько строгими должны быть критерии проверки.*/