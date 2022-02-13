package group

import (
	"database/sql"
	"errors"
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
		return fmt.Errorf("group: %w", err)
	}

	//if err := g.isValidateData(m); err != nil {
	//	return err
	//}

	//newError := model.NewError()
	// ESTE metodo de ver si existe un grupo con ese nombre es ambiguo es decir
	// necesitamos solo buscar en los grupos que esta agregado el usuario y ver si existe un nombre igual
	// de hecho mejor estos metodos de vaidar que vallan en un RELATION user algo asi
	//group, err := g.getByName(m.Name)
	//if err != nil {
	//	return fmt.Errorf("group.u.getByName(): %w", err)
	//}
	//if group.HasID() {
	//	newError.SetError(fmt.Errorf("Oops! A group with that name already exists"))
	//	newError.SetAPIMessage("¡Upps! Ya existe un grupo con ese nombre")
	//
	//	return newError
	//}

	if err := g.storage.Create(m); err != nil {
		return fmt.Errorf("group.storage.Create(): %w", err)
	}

	return nil
}

func (g Group) Update(m *model.Group) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("group: %w", err)
	}

	if err := g.storage.Update(m); err != nil {
		return fmt.Errorf("group.storage.Create(): %w", err)
	}

	return nil
}

func (g Group) UpdatePicture(m *model.Group) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("group: %w", err)
	}

	if err := g.storage.UpdatePicture(m); err != nil {
		return fmt.Errorf("group.storage.Create(): %w", err)
	}

	return nil
}

func (g Group) DeleteSoft(ID uint) error {
	if err := g.storage.DeleteSoft(ID); err != nil {
		return fmt.Errorf("User.DeleteSoft: could not delete the record %d, %w", ID, err)
	}

	return nil
}

func (g Group) Delete(ID uint) error {
	if err := g.storage.Delete(ID); err != nil {
		return fmt.Errorf("User.Delete: could not delete the record %d, %w", ID, err)
	}

	return nil
}

func (g Group) GetByID(ID uint) (model.Group, error) {
	m, err := g.storage.GetWhere(
		models.FieldsSpecification{
			models.Fields{
				{Name: "id", Value: ID},
				{Name: "deleted_at", Operator: models.IsNull},
			},
			models.SortFields{},
			models.Pagination{},
		},
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Group{}, nil
	}
	if err != nil {
		return model.Group{}, err
	}

	return m, nil
}

func (g Group) GetWhere(specification models.FieldsSpecification) (model.Group, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Group{}, fmt.Errorf("group: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.Group{}, fmt.Errorf("group: %w", err)
	}

	group, err := g.storage.GetWhere(specification)
	if err != nil {
		return model.Group{}, fmt.Errorf("group:  %w", err)
	}

	return group, nil
}

// en relation
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
