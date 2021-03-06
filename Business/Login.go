package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/data"
	entity "github.com/zcv8/YM.JinLiRead/entities"
	"github.com/zcv8/YM.JinLiRead/validation"
)

//全局的Session管理器
var Manager *common.SessionManager

//初始化函数
func init() {
	Manager, _ = common.NewManager("memory", common.AuthorizationKey, 3600)
	go Manager.GC()
}

//处理登录业务
func Login(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	ischecked := r.PostFormValue("checked")
	md5Password := common.EncryptionMD5(password)
	if !common.ValidEmail(username) && !common.ValidPhone(username) {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("Incorrect format of email or cell phone").Error(),
		})
		fmt.Fprint(wr, string(rtr))
		return
	}
	if len(password) < 6 || len(password) > 16 {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("Password length should be greater than 6 bits and less than 16 bits").Error(),
		})
		fmt.Fprint(wr, string(rtr))
		return
	}
	user, err := data.GetUser(username)
	if err == nil && md5Password == user.Password {
		tempTime := 3600
		if ischecked == "on" {
			tempTime = 7 * 24 * 3600
		}
		user.LastLoginIP = r.RemoteAddr
		user.LastLoginTime = time.Now()
		err = data.UpdateUserLastLogin(user.Id, user.LastLoginIP, user.LastLoginTime)
		if err != nil {
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status:  entity.FAILED,
				Data:    user,
				Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
				Cookie:  "",
			})
			fmt.Fprint(wr, string(rtr))
		} else {
			session := Manager.SessionStart(wr, r, int64(tempTime))
			session.Set(session.SessionID(), user.Id)
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status: entity.SUCCEED,
				Data:   user,
				Cookie: fmt.Sprintf("%s=%s;Path=/; Domain=lovemoqing.com;Max-Age=%d",
					common.AuthorizationKey, session.SessionID(), tempTime),
			})
			fmt.Fprint(wr, string(rtr))
		}
	} else {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Data:    err,
			Message: common.InterfaceUsageError.SetText("Passwords do not match").Error(),
		})
		fmt.Fprint(wr, string(rtr))
	}
}

//处理注册业务
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	md5Password := common.EncryptionMD5(password)
	if !common.ValidEmail(username) && !common.ValidPhone(username) {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("Incorrect format of email or cell phone").Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	if len(password) < 6 || len(password) > 16 {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("Password length should be greater than 6 bits and less than 16 bits").Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}

	res := validation.CaptchaVerifyHandler(w, r)
	if !res {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("Incorrect verification code").Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}

	user, err := data.InsertUser(username, md5Password)
	if err != nil {
		var Message string
		if err.Error() == "Exist" {
			Message = common.InterfaceUsageError.SetText("The user already exists. Please re-register").Error()
		} else {
			Message = common.ApplicationInternalError.SetOrginalError(err).Error()
		}
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: Message,
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	session := Manager.SessionStart(w, r, int64(3600))
	user.LastLoginIP = r.RemoteAddr
	user.LastLoginTime = time.Now()
	err = data.UpdateUserLastLogin(user.Id, user.LastLoginIP, user.LastLoginTime)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
			Cookie:  "",
		})
		fmt.Fprint(w, string(rtr))
	} else {
		session.Set(session.SessionID(), user.Id)
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status: entity.SUCCEED,
			Data:   user,
			Cookie: fmt.Sprintf("%s=%s;Path=/; Domain=lovemoqing.com;Max-Age=%d",
				common.AuthorizationKey, session.SessionID(), 3600),
		})
		fmt.Fprint(w, string(rtr))
	}
	return
}

//处理登出业务
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Manager.SessionDestroy(w, r)
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status:  entity.SUCCEED,
		Message: common.InterfaceUsageError.SetText("INVALID_SESSION").Error(),
		Cookie: fmt.Sprintf("%s='';Path=/; Domain=lovemoqing.com;Max-Age=-1",
			common.AuthorizationKey),
	})
	fmt.Fprint(w, string(rtr))
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
func ValidLoginStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, _ := Manager.SessionRead(w, r)
	if session == nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.InterfaceUsageError.SetText("INVALID_SESSION").Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	userId := session.Get(session.SessionID()).(int)
	user, err := data.GetUserById(userId)
	if err != nil {
		rtr, _ := json.Marshal(&entity.ResponseStatus{
			Status:  entity.FAILED,
			Message: common.ApplicationInternalError.SetOrginalError(err).Error(),
		})
		fmt.Fprint(w, string(rtr))
		return
	}
	rtr, _ := json.Marshal(&entity.ResponseStatus{
		Status: entity.SUCCEED,
		Data:   user,
	})
	fmt.Fprint(w, string(rtr))
}
