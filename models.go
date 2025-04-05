package main

import (
	"time"

	"github.com/dd3vahmad/rss/db"
	"github.com/google/uuid"
)

type User struct {
	ID	uuid.UUID `json:"id"`
	Name	string    `json:"name"`
	CreatedAt	time.Time    `json:"created_at"`
	UpdatedAt	time.Time    `json:"updated_at"`
}

func databaseUserToUser (dbUser db.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}
}