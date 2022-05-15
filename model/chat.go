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
	ContactID uint      `json:"contact_id,omitempty"` // el contact es null porque si queremos crear grupo que haremos con  el contact?
	GroupID   uint      `json:"group_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Chats []Chat
