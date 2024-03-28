package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rohan03122001/RSSBlog/internal/database"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Error Parsing JSON", err))
		return
	}

	

	user, err := apicfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	respondWithJSON(w, 200, struct{}{})
}
