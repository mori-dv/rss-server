package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	Update := Update{}
	err := decoder.Decode(&Update)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	}
	fmt.Println(Update)
	responseWithJSON(w, http.StatusOK, Update)
}
