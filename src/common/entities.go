package common

//响应的实体
type ReturnStatus struct {
	Status  string
	Data    interface{}
	ErrCode string
}
