package message

import (
	"database/sql"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"

	"github.com/AJRDRGZ/db-query-builder/models"

	"github.com/MikelSot/clonewhatsapp/model"
)

const table = "messages"

var fields = []string{
	"message",
	"type",
	"start",
	"user_id",
	"chat_id",
}

var constraints = postgres.Constraints{
	"messages_user_id_fk": model.ErrMessagesUserIDFK,
	"messages_chat_id_fk": model.ErrMessagesChatIDFK,
}

var (
	psqlInsert      = postgres.BuildSQLInsert(table, fields)
	psqlUpdateStart = "UPDATE " + table + " SET start = $1 WHERE id = $2"
	psqlDeleteSoft  = "UPDATE " + table + " SET deleted_at = now() WHERE id = $1"
	psqlDelete      = "DELETE FROM " + table + " WHERE id = $1"
	psqlGetAll      = postgres.BuildSQLSelect(table, fields)
)

type Message struct {
	db *sql.DB
}

func New(db *sql.DB) Message {
	return Message{db}
}

func (m2 Message) Create(m *model.Message) error {
	stmt, err := m2.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		sqlutil.StringToNull(m.Message),
		m.Type,
		m.Start,
		m.UserID,
		m.ChatID,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

func (m2 Message) UpdateStart(m *model.Message) error {
	stmt, err := m2.db.Prepare(psqlUpdateStart)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, m.Start, m.ID)
}

func (m2 Message) DeleteSoft(ID uint) error {
	stmt, err := m2.db.Prepare(psqlDeleteSoft)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (m2 Message) Delete(ID uint) error {
	stmt, err := m2.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (m2 Message) GetWhere(specification models.FieldsSpecification) (model.Message, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := m2.db.Prepare(query)
	if err != nil {
		return model.Message{}, err
	}
	defer stmt.Close()

	return m2.scanRow(stmt.QueryRow(args...))
}

func (m2 Message) scanRow(s sqlutil.RowScanner) (model.Message, error) {
	m := model.Message{}

	messageNull := sql.NullString{}
	deleteAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&messageNull,
		&m.Type,
		&m.Start,
		&m.UserID,
		&m.ChatID,
		&m.CreatedAt,
		&deleteAtNull,
	)
	if err != nil {
		return m, err
	}

	m.Message = messageNull.String
	m.DeletedAt = deleteAtNull.Time

	return m, nil
}
