package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type transact struct {
	dbConn *sqlx.DB
}

type Transact interface {
	Commit(tx *sqlx.Tx)
	Open() *sqlx.Tx
}

func NewTransact(dbConn *sqlx.DB) Transact {
	return &transact{
		dbConn: dbConn,
	}
}

func (transact *transact) Open() *sqlx.Tx {
	return transact.dbConn.MustBegin()
}

func (transact *transact) Commit(tx *sqlx.Tx) {
	err := tx.Commit()

	if err != nil {
		log.Println("Ada error", err)
		tx.Rollback()
	}
}
