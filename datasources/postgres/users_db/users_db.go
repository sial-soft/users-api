package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	Client   *sql.DB
	host     = os.Getenv("pg_users_host")
	port     = 5432
	user     = os.Getenv("pg_users_user")
	password = os.Getenv("pg_users_password")
	dbname   = os.Getenv("pg_users_database")
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
