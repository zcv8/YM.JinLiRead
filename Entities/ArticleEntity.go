package entities

import "time"

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ChannelId  int       `json:"channel"`
	Labels     string    `json:"labels"`
	Type       int       `json:"type"`
	Status     int       `json:"status"`
	FirstImage string    `json:"image"`
	Author     int       `json:"author"`
	Views      int       `json:"readcount"`
	CreateTime time.Time `json:"createtime" xorm:"created"`
	UpdateTime time.Time `json:"updatetime" xorm:"updated"`
}

//级联查询返回的结构体
type ArticleInfo struct {
	UserAdmin `xorm:"extends"`
	Channel   `xorm:"extends"`
	Article   `xorm:"extends"`
}

func (a *ArticleInfo) TableName() string {
	return "Article"
}
