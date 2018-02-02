package main

import (
	"business"
	"html/template"
	"log"
	"net/http"
	"validation"
)

func main() {
	mux := http.NewServeMux()
	staticFile := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))
	//api for create captcha
	//创建图像验证码api
	mux.HandleFunc("/api/getCaptcha", validation.GenerateCaptchaHandler)
	//api for verify captcha
	mux.HandleFunc("/api/verifyCaptcha", validation.CaptchaVerifyHandler)
	//验证登录
	mux.HandleFunc("/api/login", business.Login)
	mux.HandleFunc("/", indexHandler)

	sever := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	sever.ListenAndServe()
}

func indexHandler(wr http.ResponseWriter, r *http.Request) {
	log.Println(business.SessionManager)
	templates := []string{
		"static/templates/layout.html",
		"static/templates/index.html",
	}
	temps := template.Must(template.ParseFiles(templates...))
	temps.ExecuteTemplate(wr, "layout", "")
}
