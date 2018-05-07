package common

import (
	log "github.com/sirupsen/logrus"
)

/*
 * 日志管理 使用Logrus
		你或是你的团队写的代码
			成熟模块输出到info级别
			新写的模块 debug级别,99%的可能问题都出在这里
		第三方系统和开源框架或工具类.
			输出warn就好.
*/

/*默认写日志的方法
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
*/

func Error(text string) {
	log.Error(text)
}

func Fatal(text string) {
	log.Fatal(text)
}

func init() {
	//设置输出格式为文本格式
	log.SetFormatter(&log.TextFormatter{})
}
