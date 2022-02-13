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
}

type UseCase interface {
	Create(m *model.Message) error
	UpdateStart(m *model.Message) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetByID(ID uint) (model.Message, error)
	GetWhere(specification models.FieldsSpecification) (model.Message, error)
}
