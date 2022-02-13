package model

import "time"

type Group struct {
	ID          uint      `json:"id"`
	Picture     string    `json:"picture"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (g Group) HasID() bool { return g.ID > 0 }

type Groups []Group

func (g Groups) IsEmpty() bool {
	return len(g) < 0
}
