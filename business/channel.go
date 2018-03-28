package business

/*
 * 处理与频道标签有关的业务逻辑
 */

 import(
	 "github.com/julienschmidt/httprouter"
	 "github.com/zcv8/YM.JinLiRead/data"
	 "github.com/zcv8/YM.JinLiRead/common"
	 "encoding/json"
	 "net/http"
 )

 // 获取所有频道标签
 func GetChannels(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	 channels,err:=data.GetChannels()
	 if err!=nil{
		 json.NewEncoder(w).Encode(&common.ReturnStatus{
			Status:"failed",
			Data:err,
			ErrCode:"read failed",
		 })
		 return 
	 }
	 json.NewEncoder(w).Encode(&common.ReturnStatus{
		Status:"success",
		Data:channels,
		ErrCode:"",
	 })
}