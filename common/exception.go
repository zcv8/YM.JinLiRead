package common

import (
	"fmt"
)

/*
 * 错误管理（均实现error接口）
 * 错误主要分为2种类型：
		1.应用程序内部异常，由于程序BUG导致的异常
			* 对外接口显示 “An error occurred, Please contact the administrator”
			* 应用程序内部日志记录异常具体信息，记录 Error 级别日志
		2.接口使用不正确导致的异常，比如文件格式错误等
			* 对外接口显示具体的错误信息以此来提示用户按照错误信息进行调整
			* 应用程序内部不错任何处理
*/

var (
	ApplicationInternalError = applicationInternalError{}
	InterfaceUsageError      = interfaceUsageError{}
)

type baseError struct {
	Text string
}

//实现error接口
func (err *baseError) Error() string {
	return err.Text
}

//设置错误的文本
func (err *baseError) SetText(text string) *baseError {
	err.Text = text
	return err
}

type applicationInternalError struct {
	baseError
	orginalError error
}

//设置原始错误
func (err *applicationInternalError) SetOrginalError(originErr error) *applicationInternalError {
	err.orginalError = originErr
	return err
}

//实现error接口
func (err *applicationInternalError) Error() string {
	if err.Text != "" {
		Error(fmt.Sprintf("<InternalError>:%s", err.Text))
	}
	if err.orginalError != nil {
		Error(fmt.Sprintf("<OrginalError>:%s", err.orginalError.Error()))
	}
	return "An error occurred, Please contact the administrator"
}

type interfaceUsageError struct {
	baseError
}
