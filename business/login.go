package business

import (
	"encoding/json"
	"fmt"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	"github.com/zcv8/YM.JinLiRead/validation"
	"net/http"
	_"log"
)

//全局的Session管理器
var Manager *common.SessionManager

//初始化函数
func init() {
	Manager, _ = common.NewManager("memory", common.AuthorizationKey, 3600)
	go Manager.GC()
}

//处理登录业务
func Login(wr http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	ischecked := r.PostFormValue("checked")
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
	user, err := data.GetUser(username)
	if err == nil && md5Password == user.Password {
		tempTime := 0
		if ischecked == "on" {
			tempTime = 7 * 24 * 3600
		}
		session := Manager.SessionStart(wr, r, int64(tempTime))
		session.Set(session.SessionID(), user.ID)
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "success",
			Data:    user,
			ErrCode: "",
		})
		fmt.Fprint(wr, string(rtr))
	} else {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    err,
			ErrCode: "Authentication Failed",
		})
		fmt.Fprint(wr, string(rtr))
	}
}

//处理注册业务
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	md5Password := common.EncryptionMD5(password)
	if !common.ValidEmail(username) && !common.ValidPhone(username) {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "ErrorFormatter",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	if len(password) < 6 || len(password) > 16 {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "Too Long",
		})
		fmt.Fprint(w, string(rtr))
		return
	}

	res := validation.CaptchaVerifyHandler(w, r)
	if !res {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "Verify Code Error",
		})
		fmt.Fprint(w, string(rtr))
		return
	}

	user, err := data.InsertUser(username, md5Password)
	if err != nil {
		var errCode string
		if err.Error() == "Exist" {
			errCode = "Exist"
		} else {
			errCode = "Insert Failed"
		}
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    err,
			ErrCode: errCode,
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&common.ReturnStatus{
		Status:  "success",
		Data:    user,
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
	return
}

//处理登出业务
func Logout(w http.ResponseWriter, r *http.Request) {
	Manager.SessionDestroy(w, r)
}

//验证登录状态，仅供接口使用
func IsLogin(w http.ResponseWriter, r *http.Request) (session common.Session, status bool) {
	session, _ = Manager.SessionRead(w, r)
	if session == nil {
		return nil, false
	}
	return session, true
}

//验证登录状态
func ValidLoginStatus(w http.ResponseWriter, r *http.Request) {
	session, _ := Manager.SessionRead(w, r)
	if session == nil {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "INVALID_SESSION",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	userId := session.Get(session.SessionID()).(int)
	user, err1 := data.GetUserById(userId)
	if err1 != nil {
		rtr, _ := json.Marshal(&common.ReturnStatus{
			Status:  "failed",
			Data:    nil,
			ErrCode: "INVALID_SESSION",
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&common.ReturnStatus{
		Status:  "success",
		Data:    user,
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
}
