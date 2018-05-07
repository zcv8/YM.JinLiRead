package entities

//响应的实体
type ResponseStatus struct {
	Status  ResponseStatusCode
	Data    interface{}
	Message string
	Cookie  string
}
