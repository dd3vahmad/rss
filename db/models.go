// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
