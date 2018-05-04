package data

/*
 * 处理和文章相关的数据
 */

import (
	_ "encoding/json"
	"time"

	entity "github.com/zcv8/YM.JinLiRead/entities"
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
	ChannelId  int       `json:"channel"`
	Labels     string    `json:"labels"`
	Type       int       `json:"type"`
	Status     int       `json:"status"`
	FirstImage string    `json:"image"`
	CreateUser int       `json:"user"`
	ReadCount  int       `json:"readcount"`
	CreateTime time.Time `json:"createtime" xorm:"created"`
	UpdateTime time.Time `json:"updatetime" xorm:"updated"`
}

//级联查询返回的结构体
type ArticleInfo struct {
	entity.UserAdmin `xorm:"extends"`
	Channel          `xorm:"extends"`
	Article          `xorm:"extends"`
}

func (a *ArticleInfo) TableName() string {
	return "article"
}

//插入文章
func InsertArticle(title, content string, channel Channel,
	lables string, articleType int, status int, user entity.UserAdmin) (article Article, err error) {
	article = Article{
		Title:      title,
		Content:    content,
		ChannelId:  channel.ID,
		Labels:     lables,
		Type:       articleType,
		Status:     status,
		ReadCount:  0,
		CreateUser: user.Id,
	}
	_, err = Db.Insert(&article)
	return
}

//根据频道ID获取文章
func GetArticlesByChannel(pageIndex int, pageSize int,
	channelId int) (aInfos []ArticleInfo, err error) {
	aInfos = make([]ArticleInfo, 0)
	result := Db.Join("INNER", "user", "user.id=article.createuser").Join("INNER", "channel", "channel.id = article.channelid")
	if channelId != 0 {
		result = result.Where("article.channelid=?", channelId)
	}
	rows, err := result.Limit(pageSize, (pageIndex-1)*pageSize).Rows(&aInfos)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			aInfo := ArticleInfo{}
			rows.Scan(&aInfo)
			aInfos = append(aInfos, aInfo)
		}
	}
	return
}

//根据文章ID获取文章
func GetArticlesById(id int) (aInfo ArticleInfo, err error) {
	aInfo = ArticleInfo{}
	_, err = Db.Join("INNER", "user", "user.id=article.createuser").Join("INNER", "channel", "channel.id = article.channelid").
		Where("article.id=?", id).Get(&aInfo)
	return
}
