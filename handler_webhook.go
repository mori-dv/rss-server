package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getUpdatesTelegram() {
	botApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if botApiToken == "" {
		log.Println("TELEGRAM_API_TOKEN environment variable not set")
		return
	}
	log.Println("Getting Update From telegram...")
	targetUrl := fmt.Sprintf(
		"https://api.telegram.org/bot%v/getUpdates", botApiToken,
	)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   "127.0.0.1:8880",
			}),
		},
	}
	//proxyUrl, err := url.Parse("http://proxyIp:proxyPort")
	//myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	resp, err := client.Get(targetUrl)
	if err != nil {
		return
	}
	log.Println("telegram Updates response: ")
	log.Println(resp)
	var response Update
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	chatId := response.Msg.ChatDetail.Id
	//lastPosts, dbErr := apicfg.DB.GetPostsForTelUser(r.Context(), database.GetPostsForTelUserParams{
	//	TelID: int32(chatId),
	//	Limit: 10,
	//})
	//if dbErr != nil {
	//	responseWithError(w, http.StatusBadRequest, dbErr.Error())
	//	return
	//}

	SendMessageToTelegramBot(chatId, SendMessage{msg: "this is an alert for you"})

	log.Println(chatId)

}

type SendMessage struct {
	msg string
}

func SendMessageToTelegramBot(chatID int64, msg SendMessage) {
	botApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if botApiToken == "" {
		log.Println("TELEGRAM_API_TOKEN environment variable not set")
		return
	}
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%v/sendMessage?chat_id=%v&message=%v", botApiToken, chatID, msg,
	)
	response, err := http.Get(url)
	if err != nil {
		return
	}
	log.Println(response.Status)
	log.Println(response.Body)
}
