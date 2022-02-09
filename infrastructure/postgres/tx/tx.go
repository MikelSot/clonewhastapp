package tx

import (
	"database/sql"
	"github.com/MikelSot/clonewhatsapp/model"
)

type Tx struct {
	db *sql.DB
}

func New(db *sql.DB) Tx {
	return Tx{db}
}

func (t Tx) GetTx() (model.Transaction, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
