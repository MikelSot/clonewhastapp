package user

import "github.com/MikelSot/clonewhatsapp/model"

type Storage interface {
	Create(m *model.User) error
	Update(m *model.User) error
	Delete(ID uint) error
	GetAllWhere() (model.Users, error)
	GetWhere() (model.Users, error)
}
