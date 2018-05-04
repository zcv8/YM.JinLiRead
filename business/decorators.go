package business

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
	entity "github.com/zcv8/YM.JinLiRead/entities"
)

//身份识别验证器
func Authentication(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie(common.AuthorizationKey)
		isPass := false
		if err == nil && cookie.Value != "" {
			session, _ := Manager.SessionRead(w, r)
			if session != nil {
				isPass = true
			}
		}
		if isPass {
			h(w, r, ps)
		} else {
			rtr, _ := json.Marshal(&entity.ResponseStatus{
				Status:  entity.FAILED,
				Data:    nil,
				Message: common.InvalidSessionError.String(),
			})
			fmt.Fprint(w, string(rtr))
		}
	}
}
