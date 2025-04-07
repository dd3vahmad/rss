package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dd3vahmad/rss/auth"
	"github.com/dd3vahmad/rss/db"
	"github.com/google/uuid"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create user %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(&r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth: %v", err))
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), sql.NullString{ String: apiKey, Valid: true })
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get user: %v", err))
		return 
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}