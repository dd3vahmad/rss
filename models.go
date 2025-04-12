package main

import (
	"time"

	"github.com/dd3vahmad/rss/db"
	"github.com/google/uuid"
)

type User struct {
	ID	uuid.UUID `json:"id"`
	Name	string    `json:"name"`
	APIKey	string    `json:"api_key"`
	CreatedAt	time.Time    `json:"created_at"`
	UpdatedAt	time.Time    `json:"updated_at"`
}

func databaseUserToUser (dbUser db.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey.String,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}
}

type Feed struct {
	ID		uuid.UUID `json:"id"`
	Name	string    `json:"name"`
	URL		string    `json:"url"`
	UserID	uuid.UUID `json:"user_id"`
	CreatedAt	time.Time    `json:"created_at"`
	UpdatedAt	time.Time    `json:"updated_at"`
}

func databaseFeedToFeed (dbFeed db.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt.Time,
		UpdatedAt: dbFeed.UpdatedAt.Time,
	}
}