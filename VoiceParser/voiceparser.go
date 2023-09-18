package voiceparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)
const (
	APIEndpoint   = "https://api.assemblyai.com/v2/transcript"
	APIKey        = "4e113c035a624048ba5c270571662266" // Замените YOUR_ASSEMBLY_AI_API_KEY на ваш API-ключ
	PollingPeriod = time.Second * 10            // 10 секунд, измените по своему усмотрению
)
func DownloadFile(bot *tgbotapi.BotAPI, fileID, dest string) error {
	file, err := bot.GetFileDirectURL(fileID)
	if err != nil {
		return err
	}

	response, err := http.Get(file)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}
func ReturnText() (interface{}, error) {
	// Загрузите ваш аудиофайл
	fileData, err := ioutil.ReadFile("VoiceParser/storageForvoice/voice.m4a")
	if err != nil {
		panic(err)
	}
	// Отправьте запрос на транскрибацию
	req, err := http.NewRequest("POST", APIEndpoint, bytes.NewReader(fileData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("authorization", APIKey)
	req.Header.Set("content-type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response map[string]interface{}
	json.Unmarshal(body, &response)
	transcriptID := response["id"].(string)

	// Проверяйте статус транскрибации
	for {
		time.Sleep(PollingPeriod)

		checkURL := fmt.Sprintf("%s/%s", APIEndpoint, transcriptID)
		req, err := http.NewRequest("GET", checkURL, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("authorization", APIKey)

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(body, &response)

		if response["status"] == "completed" {
			
			return response["text"] , nil
		} else if response["status"] == "failed" {
			return "", err
		}
	}
	
}
func DeleteVoice(voice string){
	os.Remove(voice)
}