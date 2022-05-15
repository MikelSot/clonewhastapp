package model

// UserGroup Structure to create group administrators and obtain them
type UserGroup struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
}
