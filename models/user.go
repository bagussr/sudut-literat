package models

import (
	"time"
)

// User is a struct that defines the user model
type User struct {
	ID 	 uint   `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`	 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}