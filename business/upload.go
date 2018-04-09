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
		Data:    "https://image.baidu.com/search/detail?ct=503316480&z=0&ipn=d&word=阅读&step_word=&hs=0&pn=6&spn=0&di=179399045540&pi=0&rn=1&tn=baiduimagedetail&is=0%2C0&istype=0&ie=utf-8&oe=utf-8&in=&cl=2&lm=-1&st=undefined&cs=747489093%2C2592934563&os=3207538272%2C437205421&simid=3536026590%2C347644877&adpicid=0&lpn=0&ln=1990&fr=&fmq=1523282647402_R&fm=&ic=undefined&s=undefined&se=&sme=&tab=0&width=undefined&height=undefined&face=undefined&ist=&jit=&cg=&bdtype=0&oriquery=&objurl=http%3A%2F%2Fi1.hexunimg.cn%2F2012-12-17%2F149152372.jpg&fromurl=ippr_z2C%24qAzdH3FAzdH3Fpjvi_z%26e3Bijx7g_z%26e3Bv54AzdH3Fda8d-8d-80AzdH3F89l8cdn08_z%26e3Bip4s&gsm=0&rpstart=0&rpnum=0&islist=&querylist=",
		ErrCode: "",
	})
	fmt.Fprint(w, string(rtr))
}
