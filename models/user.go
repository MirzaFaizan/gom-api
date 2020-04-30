package models

import (
	"time"
)

// User Schema for users data storage
type User struct {
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Age        int       `json:"age"`
	Msisdn     string    `json:"msisdn"`
	Email      string    `json:"email"`
	InsertedAt time.Time `json:"inserted_at"`
	LastUpdate time.Time `json:"last_update"`
}
