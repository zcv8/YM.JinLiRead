package data

/*
 * 处理和文章相关的数据
 */

import (
	_ "encoding/json"

	entity "github.com/zcv8/YM.JinLiRead/entities"
)

//插入文章
func InsertArticle(title, content string, channel entity.Channel,
	lables string, articleType int, status int, user entity.UserAdmin) (article entity.Article, err error) {
	article = entity.Article{
		Title:      title,
		Content:    content,
		ChannelId:  channel.Id,
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
	channelId int) (aInfos []entity.ArticleInfo, err error) {
	aInfos = make([]entity.ArticleInfo, 0)
	aInfo := entity.ArticleInfo{}
	result := Db.Join("INNER", "UserAdmin", "UserAdmin.Id=Article.CreateUser").Join("INNER", "Channel", "Channel.Id = Article.ChannelId")
	if channelId != 0 {
		result = result.Where("Article.ChannelId=?", channelId)
	}
	rows, err := result.Limit(pageSize, (pageIndex-1)*pageSize).Rows(&aInfo)

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&aInfo)
			aInfos = append(aInfos, aInfo)
		}
	}
	return
}

//根据文章ID获取文章
func GetArticlesById(id int) (aInfo entity.ArticleInfo, err error) {
	aInfo = entity.ArticleInfo{}
	_, err = Db.Join("INNER", "user", "user.id=article.createuser").Join("INNER", "channel", "channel.id = article.channelid").
		Where("article.id=?", id).Get(&aInfo)
	return
}
