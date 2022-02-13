package model

import (
	"errors"
	"time"
)

var (
	ErrChatsUserIDFK  = errors.New("chat: The user(user_id) identification must be foreign")
	ErrChatsGroupIDFK = errors.New("chat: La identificación del grupo (group_id) debe ser extranjera")
)

type Chat struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	GroupID   uint      `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Chats []Chat
