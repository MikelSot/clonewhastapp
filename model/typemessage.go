package model

type TypeMessage string

const (
	TypeMessageAudio    TypeMessage = "AUDIO"
	TypeMessageText     TypeMessage = "TEXT"
	TypeMessageDocument TypeMessage = "DOCUMENT"
	TypeMessageImage    TypeMessage = "IMAGE"
	TypeMessageVideo    TypeMessage = "VIDEO"
)
