package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

func Connect() (err error) {

	var database *sqlx.DB
	if database, err = sqlx.Open("postgres", os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	db = *database
	return
}

func GrabDB() *sqlx.DB {
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
