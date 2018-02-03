package business

import (
	"common"
	"encoding/json"
	"fmt"
	"net/http"
)

//Session
type Session struct {
	Id       string
	UserId   string
	UserName string
}

//Session管理
var SessionManager = make(map[string]Session)

//处理登录业务
func Login(wr http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	md5Password := common.EncryptionMD5(password)
	if !common.ValidEmail(username) && !common.ValidPhone(username) {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "ErrorFormatter",
		})
		fmt.Fprint(wr, string(rtr))
		return
	}
	if len(password) < 6 || len(password) > 16 {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "Too Long",
		})
		fmt.Fprint(wr, string(rtr))
		return
	}

	if username == "12345678910" && md5Password == common.EncryptionMD5("123456") {
		session := Session{
			Id:       common.GetGuid(),
			UserId:   common.GetGuid(),
			UserName: username,
		}
		_cookie := &http.Cookie{
			Name:     "loginInfo",
			Value:    session.Id,
			HttpOnly: true,
		}
		http.SetCookie(wr, _cookie)
		SessionManager[session.Id] = session
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "success",
			Data:    nil,
			ErrCode: "",
		})
		fmt.Fprint(wr, string(rtr))
	} else {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "Authentication Failed",
		})
		fmt.Fprint(wr, string(rtr))
	}
}
