package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getUpdatesTelegram() {
	botApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if botApiToken == "" {
		log.Println("TELEGRAM_API_TOKEN environment variable not set")
		return
	}
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%v/getUpdates", botApiToken,
	)
	log.Println("Getting Update From telegram...")
	resp, err := http.Get(url)
	if err != nil {
		return
	}
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
