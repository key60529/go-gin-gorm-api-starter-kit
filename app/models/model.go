package models

import "time"

// Model Model definition
type Model struct {
	ID uint `json:"id"`
}

// TimeStamp TimeStamp definition
type TimeStamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDelete SoftDelete definition
type SoftDelete struct {
	DeletedAt time.Time `json:"deleted_at"`
}
