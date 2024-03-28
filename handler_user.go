package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/rohan03122001/RSSBlog/internal/database"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `name`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Error Parsing JSON", err))
		return
	}

	
	respondWithJSON(w, 200, struct{}{})
}
