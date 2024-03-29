package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responded with 5XX error", msg)
	}

	type errResponce struct{
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponce{
		Error: msg,
	})
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)

	if err != nil{
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}
	//standard value
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}