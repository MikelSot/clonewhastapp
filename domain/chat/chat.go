package chat

import (
	"github.com/MikelSot/clonewhatsapp/model"
)

type Storage interface {
	Create(m *model.Chat) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error
}

type UseCase interface {
	Create(m *model.Chat) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error
}
