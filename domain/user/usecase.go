package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/clonewhatsapp/model"
)

const (
	defaultMinLenPassword = 6
	defaultMinLenName     = 3
)

var allowedFieldsForQuery = []string{
	"id", "first_name", "last_name", "nickname", "email", "description", "confirmed_email", "created_at", "updated_at", "deleted_at",
}

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u User) Create(m *model.User) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	if err := u.isValidateData(m); err != nil {
		return err
	}

	newError := model.NewError()
	if !m.IsValidPassword(defaultMinLenPassword) {
		newError.SetError(fmt.Errorf("Oops! invalid password error"))
		newError.SetAPIMessage("¡Upps! error password no valido")

		return newError
	}

	user, err := u.GetByEmail(m.Email)
	if err != nil {
		return fmt.Errorf("user.u.GetByEmail(): %w", err)
	}
	if user.HasID() {
		newError.SetError(fmt.Errorf("Oops! There is already a user with that email"))
		newError.SetAPIMessage("¡Upps! Ya existe un usuario con ese email")

		return newError
	}

	if err = u.storage.Create(m); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	// TODO: send confirmation email

	return nil
}

func (u User) Update(m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) ResetPassword(m *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateNickname(m *model.User) error {
	// validar que el nickname sea unico

	return nil
}

func (u User) DeleteSoft(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (u User) Delete(ID uint) error {
	if err := u.Delete(ID); err != nil {
		return fmt.Errorf("User: could not delete the record %d, %w", ID, err)
	}

	return nil
}

func (u User) GetByID(ID uint) (model.User, error) {
	m, err := u.storage.GetWhere(
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
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return m, nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	m, err := u.storage.GetWhere(
		models.FieldsSpecification{
			models.Fields{
				{Name: "email", Value: email},
				{Name: "deleted_at", Operator: models.IsNull},
			},
			models.SortFields{},
			models.Pagination{},
		},
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return m, nil
}

func (u User) GetByNickname(nickname string) (model.User, error) {
	m, err := u.storage.GetWhere(
		models.FieldsSpecification{
			models.Fields{
				{Name: "nickname", Value: nickname},
				{Name: "deleted_at", Operator: models.IsNull},
			},
			models.SortFields{},
			models.Pagination{},
		},
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, nil
	}
	if err != nil {
		return model.User{}, err
	}

	return m, nil
}

func (u User) GetWhere(specification models.FieldsSpecification) (model.User, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	user, err := u.storage.GetWhere(specification)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	return user, nil
}

func (u User) GetAllWhere(specification models.FieldsSpecification) (model.Users, error) {
	if err := specification.Filters.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	groupUsers, err := u.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return groupUsers, nil
}

func (u User) isValidateData(m *model.User) error {
	m.FirstName = strings.TrimSpace(m.FirstName)
	m.LastName = strings.TrimSpace(m.LastName)
	m.Email = strings.ToLower(strings.TrimSpace(m.Email))

	newError := model.NewError()

	if m.FirstName == "" || m.LastName == "" || m.Email == "" {
		newError.SetError(fmt.Errorf("Oops! error fields must not be empty"))
		newError.SetAPIMessage("¡Upps! error los campos no deben de estar vacíos")

		return newError
	}

	if len(m.FirstName) < defaultMinLenName || len(m.LastName) < defaultMinLenName {
		newError.SetError(fmt.Errorf("Oops! error fields are too short"))
		newError.SetAPIMessage("¡Upps! error los campos son muy cortos")

		return newError
	}

	if !m.IsValidEmail() {
		newError.SetError(fmt.Errorf("Oops! invalid email error"))
		newError.SetAPIMessage("¡Upps! error email no valido")

		return newError
	}

	return nil
}
