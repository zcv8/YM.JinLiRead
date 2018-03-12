package data

/*
 * 处理和文章相关的数据
 */

import (
	_ "encoding/json"
	"time"
)

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserId     int       `json:-`
	Type       int       `json:type`
	CreateTime time.Time `json:createtime`
	ReadCount  int       `json:readcount`
}

//插入文章
func InsertArticle(title, content string, typeId, userid int) (article Article, err error) {
	article = Article{
		Title:   title,
		Content: content,
		Type:    typeId,
		UserId:  userid,
	}
	sql := `insert into articles(title,content,userid,type) Values($1,$2,$3,$4)
		 	returning id,createtime,readcount`
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	err = stmt.QueryRow(title, content, userid).Scan(&article.ID, &article.CreateTime,
		&article.ReadCount)
	return
}
