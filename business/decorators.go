package business

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/julienschmidt/httprouter"
)
//身份识别验证器
func Authentication(h httprouter.Handle)httprouter.Handle{
	return func(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
		cookie,err:= r.Cookie(common.AuthorizationKey)
		isPass:=false
		if(err==nil&&cookie.Value!=""){
			session, _ := Manager.SessionRead(w, r)
			if session != nil {
				isPass=true
			}
		}
		if(isPass){
			h(w,r,ps)
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