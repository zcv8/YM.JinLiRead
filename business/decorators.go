package business

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/zcv8/YM.JinLiRead/common"
)

//跨域请求包装器
func AccessControlAllowOrigin(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
		h(w,r)
	}
}

//身份识别验证器
func Authentication(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		cookie,err:= r.Cookie(common.AuthorizationKey)
		isPass:=false
		if(err==nil&&cookie.Value!=""){
			session, _ := Manager.SessionRead(w, r)
			if session != nil {
				isPass=true
			}
		}
		if(isPass){
			h(w,r)
		} else{
			rtr,_:=json.Marshal(&common.ReturnStatus{
				Status:  "failed",
				Data:    nil,
				ErrCode: "INVALID_SESSION",
			})
			fmt.Fprint(w,string(rtr))
		}
	}
}