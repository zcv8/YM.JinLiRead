package common

import(
	"net/http"
)

//跨域请求包装器
func AccessControlAllowOrigin(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
		h(w,r)
	}
}