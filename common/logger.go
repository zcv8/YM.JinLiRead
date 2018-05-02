package common

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

/*
 * 日志管理
 	自己写的代码, 成熟模块输出到info级别, 新写的模块 debug级别,99%的可能问题都出在这里.别人家的模块, 输出warn就好.
	这里自己,指的是你或是你的团队.
	别人家,是指系统和开源框架或工具类.
*/

//日志级别

type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	WARN
	ERROR
	FATAL
)

//重写String方法
func (level LogLevel) String() string {
	str, ok := logLevelTexts[level]
	if ok {
		return str
	}
	return "NIL"
}

var applicationDir string = "./"

var logLevelTexts = make(map[LogLevel]string)

func init() {
	applicationDir, _ = GetPath(applicationDir)
	logLevelTexts[INFO] = "INFO"
	logLevelTexts[DEBUG] = "DEBUG"
	logLevelTexts[WARN] = "WARN"
	logLevelTexts[ERROR] = "ERROR"
	logLevelTexts[FATAL] = "FATAL"
}

func Info(str string)  { writeLog(str, INFO) }
func Debug(str string) { writeLog(str, DEBUG) }
func Error(str string) { writeLog(str, ERROR) }
func Warn(str string)  { writeLog(str, WARN) }
func Fatal(str string) { writeLog(str, FATAL); os.Exit(-1) }

//默认写日志的方法
func writeLog(msg string, level LogLevel) {
	strackTraces := make([]string, 0)
	for i := 0; i < 15; i++ {
		//逐层向上获取堆栈信息
		pc, fullPath, line, ok := runtime.Caller(i)
		if !ok {
			continue
		} else {
			//替换反斜杠操作
			newApplicationDir := strings.Replace(applicationDir, "\\", "/", -1)
			if !strings.Contains(strings.ToLower(fullPath), strings.ToLower(newApplicationDir)) {
				break
			}
			lastS := strings.LastIndex(fullPath, "/")
			if lastS < 0 {
				lastS = strings.LastIndex(fullPath, "\\")
			}
			shortPath := fullPath[lastS+1:]
			funcName := runtime.FuncForPC(pc).Name()
			if strings.HasPrefix(funcName, applicationDir) {
				funcName = funcName[len(applicationDir):]
			}
			index := strings.LastIndex(funcName, ".")
			if index > 0 {
				funcName = funcName[index+1:]
			}
			strackTraces = append(strackTraces, fmt.Sprintf("%s%s:%s:%d\n", strings.Repeat("-", i+1), shortPath, funcName, line))
		}
	}
	starckTraceString := "******************StackTrace****************** \n" + strings.Join(strackTraces, "")
	log.Println(fmt.Sprintf("[%s] %s \n %s", level, msg, starckTraceString))
}
