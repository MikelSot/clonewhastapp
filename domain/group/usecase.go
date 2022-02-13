package group

import (
	"fmt"
	"strings"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/clonewhatsapp/model"
)

const (
	defaultMinLenName = 3
)

var allowedFieldsForQuery = []string{
	"id", "picture", "name", "description", "created_at", "updated_at", "deleted_at",
}

type Group struct {
	storage Storage
}

func New(storage Storage) Group {
	return Group{storage}
}

func (g Group) Create(m *model.Group) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	if err := g.isValidateData(m); err != nil {
		return err
	}

	// validar si existe otro grupo con ese mismo nombre

	if err := g.storage.Create(m); err != nil {
		return fmt.Errorf("group.storage.Create(): %w", err)
	}

	return nil
}

func (g Group) Update(m *model.Group) error {
	//TODO implement me
	panic("implement me")
}

func (g Group) UpdatePicture(m *model.Group) error {
	//TODO implement me
	panic("implement me")
}

func (g Group) DeleteSoft(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (g Group) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (g Group) GetByID(ID uint) (model.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g Group) GetWhere(specification models.FieldsSpecification) (model.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g Group) isValidateData(m *model.Group) error {
	m.Name = strings.TrimSpace(m.Name)

	newError := model.NewError()

	if m.Name == "" {
		newError.SetError(fmt.Errorf("Oops! error name cannot be empty"))
		newError.SetAPIMessage("¡Upps! error el nombre no puede estar vacio")

		return newError
	}

	if len(m.Name) < defaultMinLenName {
		newError.SetError(fmt.Errorf("Oops! error the name field is too short"))
		newError.SetAPIMessage("¡Upps! error el campo nombre son muy corto")

		return newError
	}

	return nil
}
