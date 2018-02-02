package business

import (
	"common"
	"fmt"
	"net/http"
)

//Session
type Session struct {
	Id       int
	UserId   int
	UserName string
}

//处理登录业务
func Login(wr http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	md5Password := common.EncryptionMD5(password)
	if username == "12345678910" && md5Password == common.EncryptionMD5("123456") {
		fmt.Fprint(wr, "OK")
	} else {
		fmt.Fprint(wr, "Error")
	}
}
