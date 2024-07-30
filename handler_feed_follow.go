package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mori-dv/RSS/internal/database"
	"net/http"
	"time"
)

func (apicfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	feedFollow, err := apicfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("connot create feed follows: %v", err))
		return
	}
	responseWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apicfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("connot get feeds: %v", err))
		return
	}
	responseWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(follows))

}
