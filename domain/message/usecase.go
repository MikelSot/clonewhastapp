package message

type Message struct {
	storage Storage
}

func New(storage Storage) Message {
	return Message{storage}
}
