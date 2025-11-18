package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./db/casabaldini.sqlite")
	if err != nil {
		log.Fatal(err)
	}
}
