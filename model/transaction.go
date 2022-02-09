package model

import "database/sql"

type Transaction interface {
	Commit() error
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
	Prepare(query string) (*sql.Stmt, error)
}
