package main

import (
	_ "log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zcv8/YM.JinLiRead/business"
	"github.com/zcv8/YM.JinLiRead/validation"
)

//还需要完成的事情：
//		1. 依赖注入：可参考 .net 依赖注入框架的源码实现
//		2. 全局注册拦截器：对每一次请求都可以实现请求之前和请求之后做一些操作
//      3. 数据访问ORM

func main() {

	router := httprouter.New()
	//创建图像验证码api
	router.GET("/api/getCaptcha", business.AccessControlAllowOrigin(validation.GenerateCaptchaHandler))
	//验证登录
	router.POST("/api/login", business.AccessControlAllowOrigin(business.Login))
	//注册
	router.POST("/api/register", business.AccessControlAllowOrigin(business.Register))
	//登出
	router.GET("/api/logout", business.AccessControlAllowOrigin(business.Authentication(business.Logout)))
	//验证登录状态
	router.GET("/api/validLoginStatus", business.AccessControlAllowOrigin(business.ValidLoginStatus))
	//创建文章
	router.POST("/api/article/create", business.AccessControlAllowOrigin(business.Authentication(business.CreateArticle)))
	//上传文章图片
	router.POST("/api/uploadarticleimg", business.AccessControlAllowOrigin(business.UploadArticleImage))
	//根据频道ID获取文章
	router.GET("/api/articles/:channelId", business.AccessControlAllowOrigin(business.GetArticlesByTypeId))
	//获取频道标签
	router.GET("/api/channels", business.AccessControlAllowOrigin(business.GetChannels))
	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}
	sever.ListenAndServe()
}
