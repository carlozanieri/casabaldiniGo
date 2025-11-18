package main

import (
	"database/sql"
	"fmt"
	"os"

	//_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" // Uncomment if you enable CGO and want sqlite
)

func OpenDBFromEnv() (*sql.DB, error) {
	driver := os.Getenv("DB_DRIVER") // "postgres" | "mysql" | "sqlite3"
	dsn := os.Getenv("DB_DSN")       // connection string

	if driver == "" || dsn == "" {
		return nil, fmt.Errorf("set DB_DRIVER and DB_DSN env vars")
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	// optional ping
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

//Postgres: DB_DRIVER=postgres e DB_DSN=postgres://user:pass@localhost:5432/dbname?sslmode=disable

//MySQL: DB_DRIVER=mysql e DB_DSN=user:pass@tcp(localhost:3306)/dbname?parseTime=true

//SQLite: DB_DRIVER=sqlite3 e DB_DSN=./data.db (ricorda CGO)
