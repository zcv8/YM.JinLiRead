package common

import (
	"time"
)

//用户登录的CookieName
const AuthorizationKey = "AUTHORIZATIONKEY"

//Session及其相关操作
type SessionManager struct {
	CookieName string
	Sessions   map[string]Session
}

type Session struct {
	UuId           string
	LastAccessTime time
	Data           interface{}
}
