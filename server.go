package main

import (
	"github.com/zcv8/YM.JinLiRead/business"
	"github.com/zcv8/YM.JinLiRead/validation"
	"github.com/zcv8/YM.JinLiRead/common"
	_ "log"
	"net/http"
)



func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	//创建图像验证码api
	mux.HandleFunc("/api/getCaptcha", common.AccessControlAllowOrigin(validation.GenerateCaptchaHandler))
	//验证登录
	mux.HandleFunc("/api/login", common.AccessControlAllowOrigin(business.Login))
	//注册
	mux.HandleFunc("/api/register", common.AccessControlAllowOrigin(business.Register))
	//登出
	mux.HandleFunc("/api/logout", common.AccessControlAllowOrigin(business.Logout))
	//验证登录
	mux.HandleFunc("/api/validLoginStatus", common.AccessControlAllowOrigin(business.ValidLoginStatus))

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	sever.ListenAndServe()
}
