package data

/*
 * 处理和文章相关的数据
 */

import (
	_ "encoding/json"
	"strconv"
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
	Labels     string    `json:"labels"`
	Type       int       `json:"type"`
	Status     int       `json:"status"`
	FirstImage string    `json:"image"`
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
	stmt, err := Db.(sql)
	if err != nil {
		return
	}
	err = stmt.QueryRow(title, content, channel.ID, lables, articleType, status, user.ID).Scan(&article.ID)
	return
}

//根据频道ID获取文章
func GetArticlesByChannel(pageIndex int, pageSize int,
	channelId int) (articles []Article, err error) {
	articles = make([]Article, 0)
	where := ""
	if channelId > 0 {
		where = "where article.channelid=" + strconv.Itoa(channelId)
	}
	sql := `select article.id,article.title,article.content,article.updatetime,article.labels,article.status,
	article.type,article.readcount,channel.id,channel.name ,u.id,u.username from articles AS article 
	join channels  AS channel on article.channelid = channel.id
	join users AS u on article.createuser = u.id ` + where + ` order by article.createtime desc limit $1 offset $2`
	rows, err := Db.Query(sql, pageSize, (pageIndex-1)*pageSize)
	if err != nil {
		return
	}
	for rows.Next() {
		article := Article{}
		channel := Channel{}
		user := User{}
		article.Channel = channel
		article.CreateUser = user
		rows.Scan(&article.ID, &article.Title,
			&article.Content, &article.UpdateTime, &article.Labels,
			&article.Status, &article.Type,
			&article.ReadCount,
			&article.Channel.ID, &article.Channel.Name, &article.CreateUser.ID, &article.CreateUser.UserName)
		articles = append(articles, article)
	}
	return
}

//根据文章ID获取文章
func GetArticlesById(id int) (article Article, err error) {
	article = Article{}
	channel := Channel{}
	user := User{}
	article.Channel = channel
	article.CreateUser = user
	err = Db.QueryRow(`select article.id,article.title,article.content,article.updatetime,article.labels,article.status,
		article.type,article.readcount,channel.id, channel.name ,u.id,u.username from articles AS article 
		join channels  AS channel on article.channelid = channel.id
		join users AS u on article.createuser = u.id where article.id=$1`, id).Scan(&article.ID, &article.Title,
		&article.Content, &article.UpdateTime, &article.Labels,
		&article.Status, &article.Type,
		&article.ReadCount,
		&article.Channel.ID, &article.Channel.Name, &article.CreateUser.ID, &article.CreateUser.UserName)
	if err != nil {
		return
	}
	return
}
