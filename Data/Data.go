package data

import (
	"github.com/zcv8/YM.JinLiRead/common"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var Db *xorm.Engine

//这个函数只会加载一次
func init() {
	var err error
	Db, err = xorm.NewEngine("postgres", "host=47.98.40.152 port=5432 user=postgres password=123456 dbname=JinLiRead sslmode=disable")
	if err != nil {
		common.Fatal(err.Error())
	}
	Db.ShowSQL(true)
	Db.SetMapper(core.SameMapper{})
	return
}
