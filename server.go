package main

import (
	"github.com/zcv8/YM.JinLiRead/business"
	"github.com/zcv8/YM.JinLiRead/validation"
	"html/template"
	_ "log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	//创建图像验证码api
	mux.HandleFunc("/api/getCaptcha", validation.GenerateCaptchaHandler)
	//验证登录
	mux.HandleFunc("/api/login", business.Login)
	//注册
	mux.HandleFunc("/api/register", business.Register)
	//登出
	mux.HandleFunc("/api/logout", business.Logout)
	//验证登录
	mux.HandleFunc("/api/validLoginStatus", business.ValidLoginStatus)

	mux.HandleFunc("/", indexHandler)

	//编辑文章
	//mux.HandleFunc("/api/article/edit", articleEditHandler)
	//创建文章
	//mux.HandleFunc("/api/article/create", articleCreateHandler)
	//mux.HandleFunc("/article/create", articleCreateHandler)

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	sever.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/apps/index.html")
	t.Execute(w, "")
}
