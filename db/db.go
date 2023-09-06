package db

import (
	"database/sql"
	"fmt"
	"log"
	"new-project2/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnDB() {
    conf := config.GetConfig()

	fmt.Println("port=", conf.DB_port)

    conn := "user=" + conf.DB_username + " password=" + conf.DB_password +
        " host=" + conf.DB_host + " port=" + conf.DB_port +
        " dbname=" + conf.DB_name + " sslmode=disable"

    var err error
    db, err = sql.Open("postgres", conn)
    if err != nil {
        log.Fatalf("error connecting to your database: %s", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("error pinging your database: %s", err)
    }

    fmt.Println("success connecting to your database")
}


func CreateConn() *sql.DB {
	return db
}
