package business

/*
 * 处理有关文章的业务
 */

import (
	"encoding/json"
	"fmt"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	"log"
	"net/http"
	"strconv"
)

//创建文章
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		//错误处理
		if r := recover(); r != nil {
			log.Print("Error Type:%T", r)

		}
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    "",
			ErrCode: "Insert Failed",
		})
		fmt.Fprint(w, string(rtr))
		return
	}()

	session, res := IsLogin(w, r)
	if !res {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    "",
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
			Data:    "",
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