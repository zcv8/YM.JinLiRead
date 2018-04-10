package business

/*
 * 上传文件业务处理
 */

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/common"
)

//上传文章内的图片
func UploadArticleImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	rtr, _ := json.Marshal(&common.ReturnStatus{
		Status:  "success",
		Data:    "https://sfault-avatar.b0.upaiyun.com/372/417/3724172311-5aa5f7dd16575_huge256",
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
}
