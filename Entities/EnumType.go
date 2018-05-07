package entities

//数据响应状态枚举
type ResponseStatusCode int

const (
	FAILED ResponseStatusCode = iota
	SUCCEED
)

func (code ResponseStatusCode) String() string {
	if code == FAILED {
		return "FAILED"
	}
	if code == SUCCEED {
		return "SUCCEED"
	}
	return ""
}
