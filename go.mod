module casabaldini

go 1.25.4


require (
    github.com/gorilla/mux v1.8.0
    github.com/gorilla/sessions v1.2.1
    github.com/lib/pq v1.10.7          // PostgreSQL driver
    github.com/go-sql-driver/mysql v1.7.0 // MySQL driver
    //mattn/go-sqlite3 needs CGO; only add if you want SQLite and have CGO enabled:
     github.com/mattn/go-sqlite3 v1.14.18
)