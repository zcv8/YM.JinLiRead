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
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	entity "github.com/zcv8/YM.JinLiRead/entities"
)

//创建文章
func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		//错误处理
		if r := recover(); r != nil {
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status:  entity.FAILED,
				Data:    nil,
				Message: common.InsertDataFailedError.SetText(r.(string)).String(),
			})
			fmt.Fprint(w, string(rtr))
			return
		}
	}()

	session, res := IsLogin(w, r)
	if !res {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.InvalidSessionError.String(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	userId := session.Get(session.SessionID()).(int)
	//这里可以重构，变成模型的Valid方法，或者使用工厂根据typeid或者是statusid获取对应值
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	//处理文章中的图片，替换临时路径为永久路径
	regex, err := regexp.Compile("!\\[.*?\\]\\((http[s]?://" + common.WebApiDomain + "/static/(.+?\\.(?:re png|jpg|jpeg|bmp|gif)))\\)")
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.InsertDataFailedError.SetOrginalErr(err).String(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	imageUrls := regex.FindAllStringSubmatch(content, -1)
	for _, value := range imageUrls {
		if len(value) > 2 {
			targetUrl := value[2]
			//移动文件的路径到永久保存路径
			if strings.Index(targetUrl, "temps/") == 0 {
				targetUrl = strings.Replace(targetUrl, "temps/", "", 0)
				tempDir, err := common.GetTempDir()
				if err != nil {
					rtr, _ := json.Marshal(&entity.ResponseStatus{
						Status:  entity.FAILED,
						Data:    nil,
						Message: common.InsertDataFailedError.SetOrginalErr(err).String(),
					})
					fmt.Fprint(w, string(rtr))
					return
				}
				imagePath := tempDir + targetUrl
				permDir, err := common.GetPermDir()
				if err != nil {
					rtr, _ := json.Marshal(&entity.ResponseStatus{
						Status:  entity.FAILED,
						Data:    nil,
						Message: common.InsertDataFailedError.SetOrginalErr(err).String(),
					})
					fmt.Fprint(w, string(rtr))
					return
				}
				err = common.MoveFile(imagePath, permDir+targetUrl)
				if err != nil {
					rtr, _ := json.Marshal(&entity.ResponseStatus{
						Status:  entity.FAILED,
						Data:    nil,
						Message: common.InsertDataFailedError.SetOrginalErr(err).String(),
					})
					fmt.Fprint(w, string(rtr))
					return
				}
				//将新的路径替换老的路径
				content = strings.Replace(content, targetUrl, strings.Replace(targetUrl, "temps/", "perms/", 0), 0)
			}
		}
	}
	typeId, _ := strconv.Atoi(r.PostFormValue("typeId"))
	statusId, _ := strconv.Atoi(r.PostFormValue("statusId"))
	channelId, _ := strconv.Atoi(r.PostFormValue("channelId"))
	labels := r.PostFormValue("labels")
	article, err := data.InsertArticle(title, content,
		data.Channel{ID: channelId}, labels, typeId, statusId, entity.UserAdmin{Id: userId})
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.InsertDataFailedError.SetOrginalErr(err).String(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status:  entity.SUCCEED,
		Data:    article,
		Message: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}

//根据频道ID获取文章
func GetArticlesByChannel(w http.ResponseWriter, r *http.Request, args httprouter.Params) {
	defer func() {
		if r := recover(); r != nil {
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status:  entity.FAILED,
				Data:    nil,
				Message: common.ReadDataFailedError.SetText(r.(string)).String(),
			})
			fmt.Fprint(w, string(rtr))
		}
	}()

	channelId, _ := strconv.Atoi(args.ByName("channelId"))
	pageIndex, _ := strconv.Atoi(r.FormValue("pageIndex"))
	pageSize, _ := strconv.Atoi(r.FormValue("pageSize"))
	articles, err := data.GetArticlesByChannel(pageIndex, pageSize, channelId)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.ReadDataFailedError.SetOrginalErr(err).String(),
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
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status:  entity.SUCCEED,
		Data:    articles,
		Message: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}

//根据文章ID获取文章
func GetArticlesById(w http.ResponseWriter, r *http.Request, args httprouter.Params) {
	defer func() {
		if r := recover(); r != nil {
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status:  entity.FAILED,
				Data:    nil,
				Message: common.ReadDataFailedError.SetText(r.(string)).String(),
			})
			fmt.Fprint(w, string(rtr))
		}
	}()

	id, _ := strconv.Atoi(args.ByName("Id"))
	article, err := data.GetArticlesById(id)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    nil,
			Message: common.ReadDataFailedError.SetOrginalErr(err).String(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status:  entity.SUCCEED,
		Data:    article,
		Message: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}
