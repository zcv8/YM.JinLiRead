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

	3.2.1 安装第三方依赖包	`go get -u github.com/mojocn/base64Captcha`
	
	3.2.2 `go get golang.org/x/image` 失败解决方案:
```
		mkdir -p $GOPATH/src/golang.org/x
		cd $GOPATH/src/golang.org/x
		git clone https://github.com/golang/image.git
```
> 参考：https://github.com/mojocn/base64Captcha

### 4.登录模块的逻辑 (更新于2018.02.02)
- 4.1 添加登录注册模块的页面样式的控制和验证逻辑
- 4.2 通过在内存中创建Session结构保存用户登录成功后的信息，同时保存到Cookie中，方便对于每次Http请求验证是否登录，目前暂未实现对Session过期和Cookie过期时间的处理
- 4.3 使用正则表达式验证用户输入

### 5.实现Session管理机制 (更新于2018.02.04)
- 5.1 实现了Session管理器，控制Session的添加，修改，查询，删除及其过期自动回收
- 5.2 UML图
![Session设计UML模型](https://github.com/zcv8/YM.JinLiRead/blob/master/uml/Session模型设计?raw=true)

> 参考：https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/06.0.md

### 6.重新整理项目使用Godep进行依赖管理（更新于2018.02.27）
- 6.1 在5的基础上添加Session的指定时间内过期，比如7天内免登录功能
- 6.2 使用Godep依赖管理工具管理Go项目，并调整了Go项目的文件结构
- 6.3 需要将项目签出到src/github.com/zcv8/ 目录下，然后使用 `go get github.com/tools/godep` 安装godep工具，安装成功之后就会在bin目录下出现godep的可执行文件。具体使用方式可以运行‘godep help’命令查看

### 7.使用Postgresql数据库实现对登录注册实现（更新于2018.03.06）
- 7.1 使用`github.com/lib/pq`驱动连接远程云服务器数据库Postgresql
- 7.2 实现登录和注册业务的相关逻辑

### 8.使用Delve进行Debug Go程序（更新于2018.03.07）
- 8.1 使用Delve代替log的方式，实现对程序的调试

> 参考：https://github.com/derekparker/delve?spm=a2c4e.11153959.blogcont57578.4.24c5266eNCm09t
> 中文博客：https://yq.aliyun.com/articles/57578

### 9.添加文章创建页面及其样式（更新于2018.03.09）
- 9.1 使用`editor.md`作为文章的编辑器

> 参考：https://pandao.github.io/editor.md/index.html

### 10.抽取前端展示逻辑（更新于2018.03.14）
- 10.1 整理项目为SPA程序，抽取前端业务逻辑到单独的项目
- 10.2 使用 Vue + Webpack 重构抽掉出来的前端项目





