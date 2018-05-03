package data

import (
	"time"
)

type Channel struct {
	ID         int
	Name       string
	Remark     string
	Sort       int
	CreateTime time.Time
}

//获取所有的频道标签
func GetChannels() (channels []Channel, err error) {
	channels = make([]Channel, 0)
	rows, err := Db.Rows(&channels)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			channel := Channel{}
			rows.Scan(&channel)
			channels = append(channels, channel)
		}
	}
	return
}
