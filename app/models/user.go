package models

import "time"

// User User model
type User struct {
	Model                   // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Active        bool      `json:"active"`
	LastLoginTime time.Time `json:"last_login_time"`
	TimeStamp
	SoftDelete
}
