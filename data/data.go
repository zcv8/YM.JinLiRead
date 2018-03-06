package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

//这个函数只会加载一次
func init() {
	var err error
	Db, err = sql.Open("postgres", "host=* port=5432 user=postgres password=* dbname=JinLiRead sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
