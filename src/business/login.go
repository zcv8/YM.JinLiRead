package business

import (
	"common"
	"encoding/json"
	"fmt"
	"net/http"
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
	//isChecked := r.PostFormValue("ischecked")
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
		Manager.SessionStart(wr, r)
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
