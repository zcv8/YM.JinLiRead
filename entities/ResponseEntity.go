package entities

//响应的实体
type ResponseStatus struct {
	Status  ResponseStatusCode
	Data    interface{}
	ErrCode string
	Cookie  string
}
