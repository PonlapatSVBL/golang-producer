package mysql

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
	// DB  *sql.DB
	err error
)

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	/* DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	} */
}
