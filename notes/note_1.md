### 1.安装并运行程序 (更新于2018.01.29)
- 1.1 运行 `go install server.go` 运行成功后在bin目录下生成 server 文件
- 1.2 运行 `./server` 允许程序，在浏览器中输入`localhost:8000` 网站中的错误都会在 `./server` 所在的窗口中显示

*程序中采用相对路径的形式，所以一定要使用以上的运行方式运行，要不然会出现错误："http panic :no find file xxxxxx"*

### 2.更新首页样式 (更新于2018.01.31)
- 2.1 采用 `bootstrap 4.0` 样式
- 2.2 暂时采用内嵌的Style方式调节样式,后续会根据框架本身提供的样式替换Style或者抽出来单独的文件

### 3.添加登录注册样式 (更新于2018.02.01)
- 3.1 添加登录和注册弹出样式
- 3.2 添加验证码,使用GO开源项目生成验证码

	3.2.1 安装第三方依赖包	`go get github.com/mojocn/base64Captcha`
	3.2.2 `go get golang.org/x/image` 失败解决方案:
```
		mkdir -p $GOPATH/src/golang.org/x
		cd $GOPATH/src/golang.org/x
		git clone https://github.com/golang/image.git
```
> 参考：https://studygolang.com/articles/12050
