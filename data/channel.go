package data

import (
	"time"
)

type Channel struct {
	ID         int
	Name       string
	Remark     string
	Sort       int
	CreateTime time.Time `xorm:"created"`
}

func (channel *Channel) TableName() string {
	return "channels"
}

//获取所有的频道标签
func GetChannels() (channels []Channel, err error) {
	channel := Channel{}
	channels = make([]Channel, 0)
	rows, err := Db.Rows(&channel)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&channel)
			channels = append(channels, channel)
		}
	}
	return
}
