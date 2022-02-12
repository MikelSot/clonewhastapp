package model

import (
	"errors"
	"net"
	"regexp"
	"strings"
	"time"
)

var (
	ErrUsersNicknameUK = errors.New("The nickname must be unique")
	ErrUsersEmailUK    = errors.New("The email must be unique.")
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type User struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Nickname       string    `json:"nickname"`
	Email          string    `json:"email"`
	Password       string    `json:"password,omitempty"`
	Description    string    `json:"description"`
	Picture        string    `json:"picture"`
	ConfirmedEmail bool      `json:"confirmed_email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

func (u User) HasID() bool { return u.ID > 0 }

func (u User) IsValidEmail() bool {
	if len(u.Email) < 3 || len(u.Email) > 150 {
		return false
	}

	if !emailRegex.MatchString(u.Email) {
		return false
	}

	parts := strings.Split(u.Email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}

func (u User) IsValidPassword(lenValid int) bool {
	return len(u.Password) >= lenValid
}

type Users []User

func (u Users) IsEmpty() bool {
	return len(u) < 0
}
