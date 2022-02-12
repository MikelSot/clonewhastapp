package auth

import "github.com/MikelSot/clonewhatsapp/model"

type UseCase struct {
}

type UseCaseUser interface {
	Create(m *model.User) error
	UpdateNickname(m *model.User) error
	
	GetByID(ID uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetByNickname(nickname string) (model.User, error)
}
