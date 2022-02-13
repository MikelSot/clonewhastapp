package model

import (
	"errors"
	"time"
)

var (
	ErrMessagesUserIDFK = errors.New("message: The user(user_id) identification must be foreign")
	ErrMessagesChatIDFK = errors.New("message: The chat(chat_id) identification must be foreign")
)

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
