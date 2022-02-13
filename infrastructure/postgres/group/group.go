package group

import (
	"database/sql"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"

	"github.com/MikelSot/clonewhatsapp/model"
)

const table = "groups"

var fields = []string{
	"picture",
	"name",
	"description",
}

var (
	psqlInsert        = postgres.BuildSQLInsert(table, fields)
	psqlUpdate        = postgres.BuildSQLUpdateByID(table, append(fields[:4], fields[5:]...))
	psqlUpdatePicture = "UPDATE " + table + " SET picture = $1 WHERE id = $2"
	psqlDeleteSoft    = "UPDATE " + table + " SET deleted_at = now() WHERE id = $1"
	psqlDelete        = "DELETE FROM " + table + " WHERE id = $1"
	psqlGetAll        = postgres.BuildSQLSelect(table, fields)
)

type Group struct {
	db *sql.DB
}

func New(db *sql.DB) Group {
	return Group{db}
}

func (g Group) Create(m *model.Group) error {
	stmt, err := g.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		sqlutil.StringToNull(m.Picture),
		m.Name,
		sqlutil.StringToNull(m.Description),
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		if errPsql := postgres.CheckError(err); errPsql != nil {
			return errPsql
		}

		return err
	}

	return nil
}

func (g Group) Update(m *model.Group) error {
	stmt, err := g.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = sqlutil.ExecAffectingOneRow(
		stmt,
		sqlutil.StringToNull(m.Picture),
		m.Name,
		sqlutil.StringToNull(m.Description),
		m.ID,
	)
	if err != nil {
		if errPsql := postgres.CheckError(err); errPsql != nil {
			return errPsql
		}

		return err
	}

	return nil
}

func (g Group) UpdatePicture(m *model.Group) error {
	stmt, err := g.db.Prepare(psqlUpdatePicture)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(
		stmt,
		sqlutil.StringToNull(m.Picture),
		m.ID,
	)
}

func (g Group) DeleteSoft(ID uint) error {
	stmt, err := g.db.Prepare(psqlDeleteSoft)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (g Group) Delete(ID uint) error {
	stmt, err := g.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (g Group) GetWhere(specification models.FieldsSpecification) (model.Group, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := g.db.Prepare(query)
	if err != nil {
		return model.Group{}, err
	}
	defer stmt.Close()

	return g.scanRow(stmt.QueryRow(args...))
}

func (g Group) scanRow(s sqlutil.RowScanner) (model.Group, error) {
	m := model.Group{}

	pictureNull := sql.NullString{}
	descriptionNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&pictureNull,
		&m.Name,
		&descriptionNull,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.Description = descriptionNull.String
	m.Picture = pictureNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
