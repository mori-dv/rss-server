package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mori-dv/RSS/internal/database"
	"net/http"
	"time"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("connot create user: %v", err))
		return
	}
	responseWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apicfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apicfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apicfg.DB.GetAllUsers(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
	}
	responseWithJSON(w, http.StatusOK, databaseUserToUsers(users))
}
