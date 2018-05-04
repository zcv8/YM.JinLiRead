package data

/*
 * 所有和User相关的业务
 */

import (
	"errors"
	"time"

	"github.com/zcv8/YM.JinLiRead/common"
)

type User struct {
	ID            int
	UserName      string
	Email         string
	Phone         string
	Password      string
	State         int
	LastLoginIP   string
	LastLoginTime time.Time
	CreateTime    time.Time `xorm:"created"`
	VerifyCode    string
}

//获取用户根据用户名
func GetUser(username string) (user User, err error) {
	user = User{}
	_, err = Db.Where("email=?", username).Or("phone=?", username).Get(&user)
	return
}

//根据用户根据用户Id
func GetUserById(uid int) (user User, err error) {
	user = User{}
	_, err = Db.Id(uid).Get(&user)
	return
}

//插入用户
func InsertUser(username string, password string) (user User, err error) {
	user = User{}
	user, err = GetUser(username)
	if err == nil {
		err = errors.New("Exist")
		return
	}
	user.Password = password
	if common.ValidEmail(username) {
		user.Email = username
	} else {
		user.Phone = username
	}
	_, err = Db.Insert(&user)
	return
}

//更新用户最后一次登录的信息
func UpdateUserLastLogin(uid int, ip string, loginTime time.Time) (err error) {
	user := User{}
	user.LastLoginIP = ip
	user.LastLoginTime = loginTime
	_, err = Db.Id(uid).Cols("lastloginip", "lastlogintime").Update(&user)
	return
}
