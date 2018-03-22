package main

import (
	"github.com/zcv8/YM.JinLiRead/business"
	"github.com/zcv8/YM.JinLiRead/validation"
	_ "log"
	"net/http"
)



func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	//创建图像验证码api
	mux.HandleFunc("/api/getCaptcha", business.AccessControlAllowOrigin(validation.GenerateCaptchaHandler))
	//验证登录
	mux.HandleFunc("/api/login", business.AccessControlAllowOrigin(business.Login))
	//注册
	mux.HandleFunc("/api/register", business.AccessControlAllowOrigin(business.Register))
	//登出
	mux.HandleFunc("/api/logout", business.AccessControlAllowOrigin(business.Authentication(business.Logout)))
	//验证登录
	mux.HandleFunc("/api/validLoginStatus", business.AccessControlAllowOrigin(business.ValidLoginStatus))

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	sever.ListenAndServe()
}
