package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	//"github.com/rohan03122001/RSSBlog/internal/auth"
	"github.com/rohan03122001/RSSBlog/internal/database"
)

func (apicfg *apiConfig) handlerCreatefeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing JSON %v", err))
		return
	}

	FeedFollows, err := apicfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create Feed follows  %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(FeedFollows))
}

func (apicfg *apiConfig) handlerGetfeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't Get Feed follows  %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(FeedFollows))
}

func (apicfg *apiConfig) handlerDeletefeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "{feed_followsID}")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse feed follow id  %v", err))
		return
	}

	err = apicfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse feed follow id  %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
