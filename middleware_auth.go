package main

import (
	"fmt"
	"net/http"

	"github.com/rohan03122001/RSSBlog/internal/auth"
	"github.com/rohan03122001/RSSBlog/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    ///
	return func(w http.ResponseWriter, r *http.Request){
		apikey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error %v",err))
		return
	}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil{
			respondWithError(w,400, fmt.Sprintf("Couldnt Get user %v", err))
			return
		}
			
		handler(w,r,user)
	}
}
