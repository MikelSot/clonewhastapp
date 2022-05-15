package group

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/MikelSot/clonewhatsapp/model"
)

type Storage interface {
	Create(m *model.Group) error
	Update(m *model.Group) error
	UpdatePicture(m *model.Group) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetWhere(specification models.FieldsSpecification) (model.Group, error)
	GetAllByUserID(userID uint, pag models.Pagination) (model.Groups, error)
}

type UseCase interface {
	Create(m *model.Group) error
	Update(m *model.Group) error
	UpdatePicture(m *model.Group) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetByID(ID uint) (model.Group, error)
	GetWhere(specification models.FieldsSpecification) (model.Group, error)
	GetAllByUserID(userID uint, pag models.Pagination) (model.Groups, error)
}
