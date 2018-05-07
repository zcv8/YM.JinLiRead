package data

import (
	entity "github.com/zcv8/YM.JinLiRead/entities"
)

//获取所有的频道标签
func GetChannels() (channels []entity.Channel, err error) {
	channel := entity.Channel{}
	channels = make([]entity.Channel, 0)
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
