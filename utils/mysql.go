package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	DB *sql.DB
)

func ConnectMysql() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
