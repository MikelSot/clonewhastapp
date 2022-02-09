package model

import "time"

type Message struct {
	ID        uint        `json:"id"`
	Message   string      `json:"message"`
	Type      TypeMessage `json:"type"`
	Start     bool        `json:"start"`
	UserID    uint        `json:"user_id"`
	ChatID    uint        `json:"chat_id"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt time.Time   `json:"deleted_at"`
}

type Messages []Message
