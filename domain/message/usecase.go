package message

import (
	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/MikelSot/clonewhatsapp/model"
)

type Message struct {
	storage Storage
}

func New(storage Storage) Message {
	return Message{storage}
}

func (m2 Message) Create(m *model.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) UpdateStart(m *model.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) DeleteSoft(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) Delete(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) GetWhere(specification models.FieldsSpecification) (model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) GetAllStart(pag models.Pagination) (model.Messages, error) {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) GetAllSentToUser(chatID uint, pag models.Pagination) (model.Messages, error) {
	//TODO implement me
	panic("implement me")
}

func (m2 Message) GetAllSentToGroup(groupID uint, pag models.Pagination) (model.Messages, error) {
	//TODO implement me
	panic("implement me")
}
