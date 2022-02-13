package chat

import (
	"fmt"
	"github.com/MikelSot/clonewhatsapp/model"
)

var allowedFieldsForQuery = []string{
	"id", "user_id", "group_id", "created_at", "updated_at", "deleted_at",
}

type Chat struct {
	storage Storage
}

func New(storage Storage) Chat {
	return Chat{storage}
}

func (c Chat) Create(m *model.Chat) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	// validar usuario existe (los dos en relation creo)
	// validar si el grupo existe

	if err := c.storage.Create(m); err != nil {
		return fmt.Errorf("chat.storage.Create(): %w", err)
	}

	return nil
}

func (c Chat) DeleteSoft(ID uint) error {
	if err := c.storage.DeleteSoft(ID); err != nil {
		return fmt.Errorf("User.DeleteSoft: could not delete the record %d, %w", ID, err)
	}

	return nil
}

func (c Chat) Delete(ID uint) error {
	if err := c.storage.Delete(ID); err != nil {
		return fmt.Errorf("User.Delete: could not delete the record %d, %w", ID, err)
	}

	return nil
}
