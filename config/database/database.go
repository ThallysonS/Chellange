package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var DB *sql.DB

func ConectarBanco() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=4002 dbname=professor sslmode=disable")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FecharTX() {
	if DB != nil {
		_ = DB.Close()
	}
}

func NovaTX() (*sql.Tx, error) {
	var (
		tx  *sql.Tx
		err error
	)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-time.After(time.Duration(5) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	tx, err = DB.BeginTx(ctx, nil)
	if err != nil {
		cancel()
		fmt.Println("erro")
		return nil, err
	}

	defer func() {
		cancel()
		<-time.After(100 * time.Millisecond)
	}()

	return tx, nil
}
