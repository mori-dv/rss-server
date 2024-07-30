package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mori-dv/RSS/internal/database"
	"net/http"
	"time"
)

func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Slug:      params.Slug,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("connot create feed: %v", err))
		return
	}
	responseWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apicfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("connot get feeds: %v", err))
		return
	}
	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))

}
