package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMysql(config *Configuration) error {
	sqldb, err := sql.Open("mysql", config.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	db := sqlx.NewDb(sqldb, "mysql")
	store.Store(config.SqlxKey, db)

	return nil
}

/**
* ? txFn is a function that will be called with an initialized `Transaction` object
* ? that can be used for executing statements and queries against a database.
 */
type txFn func(*sqlx.Tx) error

/**
 * ? WithTransaction() creates a new transaction and handles rollback/commit based on the
 * ? error object returned by the `txFn`.
 */
func WithTransaction(ctx context.Context, db *sqlx.DB, fn txFn) error {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
