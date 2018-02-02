package business

import (
	"common"
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
		fmt.Fprint(wr, "ErrorFormatter")
		return
	}
	if len(password) < 6 || len(password) > 16 {
		fmt.Fprint(wr, "Too Long")
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
		fmt.Fprint(wr, "OK")
	} else {
		fmt.Fprint(wr, "Error")
	}
}
