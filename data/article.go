package data

/*
 * 处理和文章相关的数据
 */

import (
	_ "encoding/json"
	"time"
)

type ArticleType int

const (
	Translation ArticleType = iota
	Original
	Reproduced
)

type ArticleStatus int

const (
	Draft ArticleStatus = iota
	Release
)

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Channel    Channel   `json:"channel"`
	Labels     string    `json:"lables"`
	Type       int       `json:"type"`
	Status     int       `json:"status"`
	CreateUser User      `json:"user"`
	ReadCount  int       `json:"readcount"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
}

//插入文章
func InsertArticle(title, content string, channel Channel,
	lables string, articleType int, status int, user User) (article Article, err error) {
	article = Article{
		Title:      title,
		Content:    content,
		Channel:    channel,
		Labels:     lables,
		Type:       articleType,
		Status:     status,
		ReadCount:  0,
		CreateUser: user,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	sql := `insert into articles(title,content,channelid,labels,type,status,createuser) values($1,$2,$3,$4,$5,$6,$7) returning id`
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	err = stmt.QueryRow(title, content, channel.ID, lables, articleType, status, user.ID).Scan(&article.ID)
	return
}

//根据类型获取文章
func GetArticles(pageIndex int, pageSize int,
	typeId int) (articles []Article, err error) {
	articles = make([]Article, 0)
	sql := "select * from articles where type=$1 limit $2 offset $3"
	rows, err := Db.Query(sql, typeId, pageSize, (pageIndex-1)*pageSize)
	if err != nil {
		return
	}
	for rows.Next() {
		article := Article{}
		rows.Scan(&article)
		articles = append(articles, article)
	}
	return
}
