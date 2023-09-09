package openai

import (
    "encoding/json"
    "log"
    "github.com/go-resty/resty/v2"
)

const (
    apiEndpoint = "https://api.openai.com/v1/chat/completions"
	apiKey = "Ключ апи от openai"
)

func MessagefromGPT(mesFromuser string) string {
    client := resty.New()
    messages := []map[string]interface{}{
		{
			"role":    "system",
			"content": "You are a helpful assistant.",
		},
		{
			"role":    "user",
			"content": mesFromuser,
		},
	}
	
	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   messages,
			"max_tokens": 150,
		}).
		Post(apiEndpoint)

    if err != nil {
        log.Printf("Error while sending request: %v", err)
        return "Sorry, I'm currently unable to process your request."
    }

    var data map[string]interface{}
    err = json.Unmarshal(response.Body(), &data)
    if err != nil {
        log.Printf("Error unmarshalling response: %v", err)
        return "Sorry, I faced an issue processing your request."
    }

    if errData, ok := data["error"].(map[string]interface{}); ok {
        errMsg, _ := errData["message"].(string)
        log.Printf("OpenAI API error: %s", errMsg)
        return "Sorry, I'm currently unavailable."
    }

    content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
    return content
}
