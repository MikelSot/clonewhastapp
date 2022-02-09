package model

import "time"

type Chat struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	GroupID   uint      `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Chats []Chat
