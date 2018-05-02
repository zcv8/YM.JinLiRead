package data

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var Db *xorm.Engine

//这个函数只会加载一次
func init() {
	var err error
	Db, err = xorm.NewEngine("postgres", "host=47.98.40.152 port=5432 user=postgres password=123456 dbname=JinLiRead sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
