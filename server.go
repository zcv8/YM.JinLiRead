package main

import (
	"fmt"
	_ "log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/business"
	"github.com/zcv8/YM.JinLiRead/common"
	"github.com/zcv8/YM.JinLiRead/validation"
)

//还需要完成的事情：
//		1. 依赖注入：可参考 .net 依赖注入框架的源码实现
//		2. 全局注册拦截器：对每一次请求都可以实现请求之前和请求之后做一些操作
//      3. 数据访问ORM

func main() {

	router := httprouter.New()
	//静态文件访问
	router.GET("/static/*fileName", staticHandler)
	//创建图像验证码api
	router.GET("/api/getCaptcha", validation.GenerateCaptchaHandler)
	//验证登录
	router.POST("/api/login", business.Login)
	//注册
	router.POST("/api/register", business.Register)
	//登出
	router.GET("/api/logout", business.Authentication(business.Logout))
	//验证登录状态
	router.GET("/api/validLoginStatus", business.ValidLoginStatus)
	//创建文章
	router.POST("/api/article/create", business.Authentication(business.CreateArticle))
	//上传文章图片
	router.POST("/api/uploadarticleimg", business.UploadArticleImage)
	//根据频道ID获取文章
	router.GET("/api/articles/:channelId", business.GetArticlesByTypeId)
	//获取频道标签
	router.GET("/api/channels", business.GetChannels)

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}
	sever.ListenAndServe()
}

var fileType = map[string]string{
	"png":  "image/png",
	"jpg":  "image/jpg",
	"jpeg": "image/jpeg",
	"gif":  "image/gif",
}

//处理静态文件
func staticHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fileName := p.ByName("fileName")
	path, err := common.GetPath("./files" + fileName)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err)
		return
	}
	bytes, err := common.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err)
		return
	}
	extendName, err := common.GetFileExtendName(fileName)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err)
		return
	}
	responseType, ok := fileType[extendName]
	if !ok {
		responseType = "text/plain"
	}
	w.Header().Add("Content-Type", responseType)
	w.Write(bytes)
}
