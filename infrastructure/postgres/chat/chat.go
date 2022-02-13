package chat

import (
	"database/sql"

	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"

	"github.com/MikelSot/clonewhatsapp/model"
)

const table = "chats"

var fields = []string{
	"user_id",
	"group_id",
}

var constraints = postgres.Constraints{
	"chats_user_id_fk":  model.ErrChatsUserIDFK,
	"chats_group_id_fk": model.ErrChatsGroupIDFK,
}

var (
	psqlInsert     = postgres.BuildSQLInsert(table, fields)
	psqlDeleteSoft = "UPDATE " + table + " SET deleted_at = now() WHERE id = $1"
	psqlDelete     = "DELETE FROM " + table + " WHERE id = $1"
)

type Chat struct {
	db *sql.DB
}

func New(db *sql.DB) Chat {
	return Chat{db}
}

func (c Chat) Create(m *model.Chat) error {
	stmt, err := c.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.UserID,
		sqlutil.Int64ToNull(int64(m.GroupID)),
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

func (c Chat) DeleteSoft(ID uint) error {
	stmt, err := c.db.Prepare(psqlDeleteSoft)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (c Chat) Delete(ID uint) error {
	stmt, err := c.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}
