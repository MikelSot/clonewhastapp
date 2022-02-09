package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Psql struct {
	db *sql.DB
}

func newPsql(config Database) (*Psql, error) {
	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Server,
		config.Port,
		config.Name,
		config.SSLMode,
	)
	dbCon, err := sql.Open("postgres", dns)

	return &Psql{db: dbCon}, err
}

func (p *Psql) GetConnection() *sql.DB {
	return p.db
}

func newPsqlConnection(conf Configuration) *sql.DB {
	psql, err := newPsql(
		newDatabase(
			conf.Database.Engine,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Server,
			conf.Database.Name,
			conf.Database.SSLMode,
			conf.Database.Port,
		),
	)
	if err != nil {
		log.Fatalf("No se puedo obtener una conexion a la bd ")
	}

	return psql.GetConnection()
}
