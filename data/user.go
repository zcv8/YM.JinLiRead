package data

/*
 * 所有和User相关的业务
 */

import (
	"errors"
	"github.com/zcv8/YM.JinLiRead/common"
	_ "log"
	"time"
)

type User struct {
	ID         int
	Email      string
	Phone      string
	Password   string
	State      int
	LastLoginIP string
	LastLoginTime time.Time
	CreateTime time.Time
	VerifyCode string
}

//获取用户根据用户名
func GetUser(username string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select ID , Email, Phone,Password,CreateTime from Users where Email = $1 Or Phone = $1", username).Scan(&user.ID, &user.Email, &user.Phone, &user.Password, &user.CreateTime)
	return
}

//根据用户根据用户Id
func GetUserById(uid int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select ID , Email, Phone,Password,CreateTime from Users where ID =$1", uid).Scan(&user.ID, &user.Email, &user.Phone, &user.Password, &user.CreateTime)
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
	stmt, errs := Db.Prepare("Insert into users(email,phone,password,createTime,state) values ($1,$2,$3,$4,$5) returning id,email,phone")
	defer stmt.Close()
	if errs != nil {
		err = errs
		return
	}
	if common.ValidEmail(username) {
		err = stmt.QueryRow(username, "NULL", password, time.Now(), 0).Scan(&user.ID, &user.Email, &user.Phone)
	} else {
		err = stmt.QueryRow("NULL", username, password, time.Now(), 0).Scan(&user.ID, &user.Email, &user.Phone)
	}
	return
}

//更新用户最后一次登录的信息
func UpdateUserLastLogin(uid int,ip string,loginTime time.Time)(err error){
	defer func(){
		if r:=recover();r!=nil{
			err= errors.New("update error")
		}
	}()
	stmt,err:= Db.Prepare("UPDATE users SET lastloginip=$2,lastlogintime=$3 WHERE id = $1")
	stmt.QueryRow(uid, ip,time.Now())
	return
}
