package common

import (
	"fmt"
	"log"
)

/*
 * 错误管理
 */

//错误的文本枚举类型
type errorText string

const (
	IfWriteErrLog             bool      = true //默认设置为错误书写日志
	defaultErrText            string    = "error is not registered"
	InvalidSessionError       errorText = "innerInvalidSessionError"
	InvalidFormatterError     errorText = "innerInvalidFormatterError"
	StringTooLongError        errorText = "innerStringTooLongError"
	AuthenticationFailedError errorText = "innerAuthenticationFailedError"
	VerificationCodeError     errorText = "innerVerificationCodeError"
	UpdateDataFailedError     errorText = "innerUpdateDataFailedError"
	InsertDataFailedError     errorText = "innerInsertDataFailedError"
	ExistingDataError         errorText = "innerExistingDataError"
	ReadDataFailedError       errorText = "innerReadDataFailedError"
)

//重写String方法 返回错误代码
func (errText errorText) String() string {
	v, ok := errInfos[errText]
	if ok {
		//写日志
		if IfWriteErrLog {
			log.Println(fmt.Sprintf("[Error:%s]:%s", v.GetCode(), v.GetText()))
			log.Println(fmt.Sprintf("[InnerError]:%s", v.GetOrginalErr()))
		}
		return v.GetCode()
	}
	return defaultErrText
}

//包装原始错误
func (errText errorText) SetOrginalErr(err error) errorText {
	v, ok := errInfos[errText]
	if ok {
		v.SetOrginalErr(err)
	}
	return errText
}

//设置自定义错误文本
func (errText errorText) SetText(txt string) errorText {
	v, ok := errInfos[errText]
	if ok {
		v.SetText(txt)
	}
	return errText
}

//错误需要实现的接口方法
type errInterface interface {
	GetText() string
	GetCode() string
	GetOrginalErr() error
	SetOrginalErr(err error)
	SetText(txt string)
}

//内部错误：基础信息
type errInfo struct {
	text         string //错误代码的描述信息
	code         string //错误代码
	orginalError error  //原始错误
}

//获取错误代码文本
func (err *errInfo) GetText() string {
	return err.text
}

//设置错误代码文本
func (err *errInfo) SetText(txt string) {
	err.text = txt
}

func (err *errInfo) GetOrginalErr() error {
	return err.orginalError
}

//获取错误代码
func (err *errInfo) GetCode() string {
	return err.code
}

//设置内部异常
func (err *errInfo) SetOrginalErr(e error) {
	err.orginalError = e
}

//内部错误：未验证Session
type innerInvalidSessionError struct{ errInfo }

//内部错误：输入字符串格式验证未通过
type innerInvalidFormatterError struct{ errInfo }

//内部错误：字符串太长超出限制
type innerStringTooLongError struct{ errInfo }

//内部错误：权限验证失败
type innerAuthenticationFailedError struct{ errInfo }

//内部错误：更新数据失败
type innerUpdateDataFailedError struct{ errInfo }

//内部错误：验证码验证失败
type innerVerificationCodeError struct{ errInfo }

//内部错误：插入数据失败
type innerInsertDataFailedError struct{ errInfo }

//内部错误：数据已经存在
type innerExistingDataError struct{ errInfo }

//内部错误：读取数据失败
type innerReadDataFailedError struct{ errInfo }

//全局错误管理器，根据错误文本映射到内部错误
var errInfos map[errorText]errInterface

//初始化函数，注册内部错误到全局错误管理器
func init() {
	errInfos[InvalidSessionError] = &innerInvalidSessionError{
		errInfo: errInfo{
			text: "Invalid session",
			code: "InvalidSession",
		},
	}
	errInfos[InvalidFormatterError] = &innerInvalidFormatterError{
		errInfo: errInfo{
			text: "Incorrect input string format",
			code: "InvalidFormatter",
		},
	}
	errInfos[StringTooLongError] = &innerStringTooLongError{
		errInfo: errInfo{
			text: "Incorrect input string format",
			code: "StringTooLong",
		},
	}
	errInfos[AuthenticationFailedError] = &innerAuthenticationFailedError{
		errInfo: errInfo{
			text: "Permission verification failed",
			code: "AuthenticationFailed",
		},
	}
	errInfos[UpdateDataFailedError] = &innerUpdateDataFailedError{
		errInfo: errInfo{
			text: "Failed to update data",
			code: "UpdateDataFailed",
		},
	}
	errInfos[VerificationCodeError] = &innerVerificationCodeError{
		errInfo: errInfo{
			text: "Verification code verification failed",
			code: "VerificationCodeError",
		},
	}
	errInfos[InsertDataFailedError] = &innerInsertDataFailedError{
		errInfo: errInfo{
			text: "Failed to insert data",
			code: "InsertDataFailed",
		},
	}
	errInfos[ExistingDataError] = &innerExistingDataError{
		errInfo: errInfo{
			text: "There is an existing data error",
			code: "ExistingData",
		},
	}
	errInfos[ReadDataFailedError] = &innerReadDataFailedError{
		errInfo: errInfo{
			text: "Failed to read data",
			code: "ReadDataFailed",
		},
	}
}
