package database

import (
    "os"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/emikohmann/url-shortener/src/api/config"
)

const (
    infDatabaseConnected  = "Database connected"
    errConnectingDatabase = "Error connecting database"
)

var (
    username = os.Getenv("DB_USERNAME")
    password = os.Getenv("DB_PASSWORD")
    host     = os.Getenv("DB_HOST")
    schema   = os.Getenv("DB_SCHEMA")

    Client *sql.DB
)

func Connect() {
    connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

    var err error
    if Client, err = sql.Open("mysql", connection); err != nil {
        config.Logger.Println(errConnectingDatabase, err)
        panic(err)
    }
    
    config.Logger.Println(infDatabaseConnected)
}
