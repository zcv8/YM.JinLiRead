package entities

import "time"

//用户登录注册相关实体
type UserAdmin struct {
	Id              int
	UserName        string
	Email           string
	VerifyEmailCode string `xorm:"<-" json:"-"`
	Phone           string
	VerifyPhoneCode string `xorm:"<-" json:"-"`
	Password        string
	State           int
	LastLoginIP     string
	LastLoginTime   time.Time
	CreateTime      time.Time `xorm:"created"`
	UpdateTime      time.Time `xorm:"updated"`
}
