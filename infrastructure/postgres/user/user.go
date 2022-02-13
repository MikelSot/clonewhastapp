package user

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"
	"golang.org/x/crypto/bcrypt"

	"github.com/MikelSot/clonewhatsapp/model"
)

const table = "users"

var fields = []string{
	"first_name",
	"last_name",
	"nickname",
	"email",
	"password",
	"description",
	"picture",
	"confirmed_email",
}

var constraints = postgres.Constraints{
	"users_email_uk":    model.ErrUsersEmailUK,
	"users_nickname_uk": model.ErrUsersNicknameUK,
}

var (
	psqlInsert         = postgres.BuildSQLInsert(table, fields)
	psqlUpdate         = postgres.BuildSQLUpdateByID(table, append(fields[:4], fields[5:]...))
	psqlResetPassword  = "UPDATE users SET password = $1 WHERE id = $2"
	psqlUpdateNickname = "UPDATE users SET nickname = $1 WHERE id = $2"
	psqlDeleteSoft     = "UPDATE " + table + " SET deleted_at = now() WHERE id = $1"
	psqlDelete         = "DELETE FROM " + table + " WHERE id = $1"
	psqlGetAll         = postgres.BuildSQLSelect(table, append(fields[:4], fields[5:]...))
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) Create(m *model.User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generando el hash del password: %v", err)
	}

	stmt, err := u.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.FirstName,
		m.LastName,
		sqlutil.StringToNull(m.Nickname),
		m.Email,
		pass,
		sqlutil.StringToNull(m.Description),
		sqlutil.StringToNull(m.Picture),
		m.ConfirmedEmail,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

func (u User) Update(m *model.User) error {
	stmt, err := u.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = sqlutil.ExecAffectingOneRow(
		stmt,
		m.FirstName,
		m.LastName,
		sqlutil.StringToNull(m.Nickname),
		m.Email,
		sqlutil.StringToNull(m.Description),
		sqlutil.StringToNull(m.Picture),
		m.ConfirmedEmail,
		m.ID,
	)
	if err != nil {
		return postgres.CheckConstraint(constraints, err)
	}

	return nil
}

func (u User) ResetPassword(m *model.User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generando el hash del password: %v", err)
	}

	stmt, err := u.db.Prepare(psqlResetPassword)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, pass, m.ID)
}

func (u User) UpdateNickname(m *model.User) error {
	stmt, err := u.db.Prepare(psqlUpdateNickname)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, m.Password, m.ID)
}

func (u User) DeleteSoft(ID uint) error {
	stmt, err := u.db.Prepare(psqlDeleteSoft)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (u User) Delete(ID uint) error {
	stmt, err := u.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return sqlutil.ExecAffectingOneRow(stmt, ID)
}

func (u User) GetWhere(specification models.FieldsSpecification) (model.User, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return model.User{}, err
	}
	defer stmt.Close()

	return u.scanRow(stmt.QueryRow(args...))
}

func (u User) GetAllWhere(specification models.FieldsSpecification) (model.Users, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)
	query += " " + postgres.BuildSQLPagination(specification.Pagination)

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s sqlutil.RowScanner) (model.User, error) {
	m := model.User{}
	nicknameNull := sql.NullString{}
	descriptionNull := sql.NullString{}
	pictureNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.FirstName,
		&m.LastName,
		&nicknameNull,
		&m.Email,
		&descriptionNull,
		&pictureNull,
		&m.ConfirmedEmail,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.Nickname = nicknameNull.String
	m.Description = descriptionNull.String
	m.Picture = pictureNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
