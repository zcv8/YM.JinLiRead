package business

/*
 * 处理文章有关的业务逻辑
 */

import (
	"encoding/json"
	"fmt"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	"github.com/julienschmidt/httprouter"
	_"log"
	"net/http"
	"strconv"
)

//创建文章
func CreateArticle(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
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
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	typeId, _ := strconv.Atoi(r.PostFormValue("typeId"))
	article, err := data.InsertArticle(title, content, typeId, userId)
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
