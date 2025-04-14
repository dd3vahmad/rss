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

func (apiCfg *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := apiCfg.DB.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get users: %v", err))
		return
	}

	usersList := make([]User, len(users))
	for i, user := range users {
		usersList[i] = databaseUserToUser(user)
	}
	if len(usersList) == 0 {
		respondWithError(w, 404, "No users found")
		return
	}

	respondWithJSON(w, 200, usersList)
}

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID: uuid.New(),
		Name: params.Name,
		Url: params.URL,
		UserID: params.UserID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create feed because: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}