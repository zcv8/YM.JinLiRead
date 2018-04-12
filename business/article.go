package business

/*
 * 处理文章有关的业务逻辑
 */

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
)

//创建文章
func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		//错误处理
		if r := recover(); r != nil {
			rtr, _ := json.Marshal(&common.ReturnStatus{
				Status:  "failed",
				Data:    r,
				ErrCode: "Insert Failed",
			})
			fmt.Fprint(w, string(rtr))
			return
		}
	}()

	session, res := IsLogin(w, r)
	if !res {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    res,
			ErrCode: "INVALID_SESSION",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	userId := session.Get(session.SessionID()).(int)
	//这里可以重构，变成模型的Valid方法，或者使用工厂根据typeid或者是statusid获取对应值
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	typeId, _ := strconv.Atoi(r.PostFormValue("typeId"))
	statusId, _ := strconv.Atoi(r.PostFormValue("statusId"))
	channelId, _ := strconv.Atoi(r.PostFormValue("channelId"))
	labels := r.PostFormValue("labels")
	article, err := data.InsertArticle(title, content,
		data.Channel{ID: channelId}, labels, typeId, statusId, data.User{ID: userId})
	if err != nil {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    err,
			ErrCode: "Insert Failed",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&common.ReturnStatus{
		Status:  "success",
		Data:    article,
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}

//根据频道ID获取文章
func GetArticlesByTypeId(w http.ResponseWriter, r *http.Request, args httprouter.Params) {
	defer func() {
		if r := recover(); r != nil {
			rtr, _ := json.Marshal(&common.ReturnStatus{
				Status:  "failed",
				Data:    r,
				ErrCode: "An error occurred",
			})
			fmt.Fprint(w, string(rtr))
		}
	}()

	channelId, _ := strconv.Atoi(args.ByName("channelId"))
	pageIndex, _ := strconv.Atoi(r.FormValue("pageIndex"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	articles, err := data.GetArticles(pageIndex, pageSize, channelId)
	if err != nil {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    err,
			ErrCode: "An error occurred",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	for index, value := range articles {
		//正则匹配替换图片为"[图片]"字样
		regex, err := regexp.Compile("!\\[.*?\\]\\((http[s]?://.+?\\.(png|jpg|jpeg|bmp|gif))\\)")
		if err != nil {
			articles[index].Content = "内容获取失败"
		} else {
			articles[index].Content = regex.ReplaceAllString(value.Content, "[图片]")
		}
		vlen := len(articles[index].Content)
		if vlen > 200 {
			articles[index].Content = articles[index].Content[0:200] + "..."
		} else {
			articles[index].Content = articles[index].Content + "..."
		}
		//处理文章首图
		matchs := regex.FindSubmatch([]byte(value.Content))
		if len(matchs) > 1 {
			articles[index].FirstImage = string(matchs[1])
		}
	}
	rtr, _ := json.Marshal(&common.ReturnStatus{
		Status:  "success",
		Data:    articles,
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}
