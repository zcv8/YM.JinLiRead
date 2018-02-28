package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
