package entities

import (
	"time"
)

type Channel struct {
	Id         int
	Name       string
	Remark     string
	CreateUser int       `json:"-"`
	CreateTime time.Time `xorm:"created" json:"-"`
	UpdateTime time.Time `xorm:"updated" json:"-"`
}
