package model

type UserMessage struct {
	User     User     `json:"user"`
	Messages Messages `json:"messages"`
}

type UserMessages []UserMessage

type TotalForUser struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	Contacts uint `json:"contacts"`
	Groups   uint `json:"group"`
}
