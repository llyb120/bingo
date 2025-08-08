package sqly

import (
	"database/sql"
	"database/sql/driver"
)

type DBLike interface {
	Prepare(query string) (*sql.Stmt, error)
	Driver() driver.Driver
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Tx struct {
	Db *sql.DB
	Tx *sql.Tx
}

func (t *Tx) Prepare(query string) (*sql.Stmt, error) {
	return t.Tx.Prepare(query)
}

func (t *Tx) Driver() driver.Driver {
	return t.Db.Driver()
}

func (t *Tx) Commit() error {
	return t.Tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}

func (t *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.Tx.Exec(query, args...)
}

func NewTx(db *sql.DB) (*Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{Db: db, Tx: tx}, nil
}
