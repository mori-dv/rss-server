// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Slug          string
	UserID        uuid.UUID
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type Post struct {
	ID          uuid.UUID
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	Url         string
	Description sql.NullString
	FeedID      uuid.UUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
	IsAdmin   string
	TelID     sql.NullString
}
