package model

import "time"

type User struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Nickname       string    `json:"nickname"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Description    string    `json:"description"`
	Picture        string    `json:"picture"`
	ConfirmedEmail bool      `json:"confirmed_email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type Users []User
