package message

import (
	"database/sql"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"

	"github.com/MikelSot/clonewhatsapp/model"
)

const table = "messages"

var fields = []string{
	"message",
	"type",
	"start",
	"user_id",
	"chat_id",
	"group_id",
}

var constraints = postgres.Constraints{
	"messages_user_id_fk": model.ErrMessagesUserIDFK,
	"messages_chat_id_fk": model.ErrMessagesChatIDFK,
}

var (
	psqlInsert        = postgres.BuildSQLInsert(table, fields)
	psqlUpdateStart   = "UPDATE " + table + " SET start = $1 WHERE id = $2"
	psqlDeleteSoft    = "UPDATE " + table + " SET deleted_at = now() WHERE id = $1"
	psqlDelete        = "DELETE FROM " + table + " WHERE id = $1"
	psqlGetAll        = postgres.BuildSQLSelect(table, fields)
	psqlGetAllStart   = psqlGetAll + " WHERE start = true"
	psqlGetSentToUser = psqlGetAll + `AS m WHERE chat_id = $1 AND m.deleted_at IS NULL
					  ORDER BY m.created_at DESC , m.id DESC`
	psqlGetSentToGroup = psqlGetAll + `AS m WHERE group_id = $1 AND m.deleted_at IS NULL
					   ORDER BY m.created_at DESC , m.id DESC;`
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
		sqlutil.Int64ToNull(int64(m.GroupID)),
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

func (m2 Message) GetAllStart(pag models.Pagination) (model.Messages, error) {
	query := psqlGetAllStart + " " + postgres.BuildSQLPagination(pag)
	stmt, err := m2.db.Prepare(query)
	if err != nil {
		return model.Messages{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return model.Messages{}, err
	}
	defer rows.Close()

	ms := make(model.Messages, 0)
	for rows.Next() {
		m, err := m2.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (m2 Message) GetAllSentToUser(chatID uint, pag models.Pagination) (model.Messages, error) {
	query := psqlGetSentToUser + " " + postgres.BuildSQLPagination(pag)
	stmt, err := m2.db.Prepare(query)
	if err != nil {
		return model.Messages{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(chatID)
	if err != nil {
		return model.Messages{}, err
	}
	defer rows.Close()

	ms := make(model.Messages, 0)
	for rows.Next() {
		m, err := m2.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (m2 Message) GetAllSentToGroup(groupID uint, pag models.Pagination) (model.Messages, error) {
	query := psqlGetSentToGroup + " " + postgres.BuildSQLPagination(pag)
	stmt, err := m2.db.Prepare(query)
	if err != nil {
		return model.Messages{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return model.Messages{}, err
	}
	defer rows.Close()

	ms := make(model.Messages, 0)
	for rows.Next() {
		m, err := m2.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (m2 Message) scanRow(s sqlutil.RowScanner) (model.Message, error) {
	m := model.Message{}

	messageNull := sql.NullString{}
	groupIDNull := sql.NullInt64{}
	deleteAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&messageNull,
		&m.Type,
		&m.Start,
		&m.UserID,
		&m.ChatID,
		&groupIDNull,
		&m.CreatedAt,
		&deleteAtNull,
	)
	if err != nil {
		return m, err
	}

	m.Message = messageNull.String
	m.GroupID = uint(groupIDNull.Int64)
	m.DeletedAt = deleteAtNull.Time

	return m, nil
}
