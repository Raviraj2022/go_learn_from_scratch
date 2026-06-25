package config

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    var err error

    DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("DB connection error:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("DB not reachable:", err)
    }

    log.Println("Connected to Neon DB 🚀")
}