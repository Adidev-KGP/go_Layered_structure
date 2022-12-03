package db

import (
    "database/sql"
    "fmt"
    "log"

    "os"

    _ "github.com/lib/pq"
)

var (
    Client *sql.DB

    host     = os.Getenv("DB_HOST")
    port     = os.Getenv("DB_PORT")
    user     = os.Getenv("DB_USERNAME")
    password = os.Getenv("DB_PASSWORD")
    dbname   = os.Getenv("DB_NAME")
)

func init() {
    fmt.Println("Connecting to database")
    postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    var err error
    Client, err = sql.Open("postgres", postgresqlDbInfo)
    if err != nil {
        panic(err)
    }

    log.Println("database successfully configured")
}