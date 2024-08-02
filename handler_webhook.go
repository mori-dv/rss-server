package main

import (
	"encoding/json"
	"fmt"
	"github.com/mori-dv/RSS/internal/database"
	"log"
	"net/http"
	"os"
	"time"
)

func (apicfg *apiConfig) webhookHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	Update := Update{}
	err := decoder.Decode(&Update)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	}
	lastPosts, dbErr := apicfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if dbErr != nil {
		responseWithError(w, http.StatusBadRequest, dbErr.Error())
	}

	fmt.Println(lastPosts)

	if dbErr != nil {
		return
	}
	chatId := Update.Msg.ChatDetail.Id
	if chatId != 0 {
		responseWithError(w, http.StatusBadRequest, "You can't send messages to this chat")
	}
	responseWithJSON(w, http.StatusOK, Update)
}

type SendMessage struct {
}

func SendMessageToTelegramBot(chatID int64, msg SendMessage) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	botApiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if botApiToken == "" {
		log.Println("TELEGRAM_BOT_API_TOKEN environment variable not set")
		return
	}
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%v/sendMessage?chat_id=%v&message=%v", botApiToken, chatID, msg,
	)
	response, err := client.Get(url)
	if err != nil {
		return
	}
	log.Println(response.Status)
	log.Println(response.Body)
}
