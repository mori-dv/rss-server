package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error ", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errorResponse{Error: message})

}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write header: %v\n", err)
		return
	}

}
