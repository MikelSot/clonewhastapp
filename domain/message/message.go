package message

import (
	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/clonewhatsapp/model"
)

type Storage interface {
	Create(m *model.Message) error
	UpdateStart(m *model.Message) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Message, error)
	GetAllStart(pag models.Pagination) (model.Messages, error)
	GetAllSentToUser(chatID uint, pag models.Pagination) (model.Messages, error)
	GetAllSentToGroup(groupID uint, pag models.Pagination) (model.Messages, error)
}

type UseCase interface {
	Create(m *model.Message) error
	UpdateStart(m *model.Message) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Message, error)
	GetAllStart(pag models.Pagination) (model.Messages, error)
	GetAllSentToUser(chatID uint, pag models.Pagination) (model.Messages, error)
	GetAllSentToGroup(groupID uint, pag models.Pagination) (model.Messages, error)
}
