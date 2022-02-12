package user

import (
	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/clonewhatsapp/model"
)

type Storage interface {
	Create(m *model.User) error
	Update(m *model.User) error
	ResetPassword(m *model.User) error
	UpdateNickname(m *model.User) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetAllWhere(specification models.FieldsSpecification) (model.Users, error)
	GetWhere(specification models.FieldsSpecification) (model.User, error)
}

type UseCase interface {
	Create(m *model.User) error
	Update(m *model.User) error
	ResetPassword(m *model.User) error
	UpdateNickname(m *model.User) error
	DeleteSoft(ID uint) error
	Delete(ID uint) error

	GetByID(ID uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetByNickname(nickname string) (model.User, error)
	GetWhere(specification models.FieldsSpecification) (model.User, error)
	GetAllWhere(specification models.FieldsSpecification) (model.Users, error)
}
