package data

import(
	"time"
)

type Channel struct{
	ID int
	Name string
	Remark string
	Sort int
	CreateTime time.Time
}

//获取所有的频道标签
func GetChannels()(channels []Channel,err error){
	channels=make([]Channel,0)
	sql:="SELECT id,name,remark,sort,createtime FROM channels ORDER BY sort"
	rows,err := Db.Query(sql)
	defer rows.Close()
	if err==nil{
		for rows.Next(){
			channel:=Channel{}
			err=rows.Scan(&channel.ID,&channel.Name,&channel.Remark,&channel.Sort,&channel.CreateTime)
			if err!=nil{
				return 
			}
			channels=append(channels,channel)
		}
	}
	return
}