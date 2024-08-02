package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	webhook := Webhook{}
	err := decoder.Decode(&webhook)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	}
	fmt.Println(webhook)
}
