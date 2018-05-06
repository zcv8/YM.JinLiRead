### 1.安装并运行程序 (更新于2018.01.29)
- 1.1 运行 `go install server.go` 运行成功后在bin目录下生成 server 文件
- 1.2 运行 `./server` 允许程序，在浏览器中输入`localhost:8000` 网站中的错误都会在 `./server` 所在的窗口中显示

*程序中采用相对路径的形式，所以一定要使用以上的运行方式运行，要不然会出现错误："http panic :no find file xxxxxx"*

### 2.更新首页样式 (更新于2018.01.31)
- 2.1 采用 `bootstrap 4.0` 样式
- 2.2 暂时采用内嵌的Style方式调节样式,后续会根据框架本身提供的样式替换Style或者抽出来单独的文件

### 3.添加验证码 (更新于2018.02.01)
- 3.1 添加验证码,使用GO开源项目生成验证码

	3.1.1 安装第三方依赖包	`go get -u github.com/mojocn/base64Captcha`
	
	3.1.2 `go get golang.org/x/image` 失败解决方案:
```
		mkdir -p $GOPATH/src/golang.org/x
		cd $GOPATH/src/golang.org/x
		git clone https://github.com/golang/image.git
```
> 参考：https://github.com/mojocn/base64Captcha

### 4.实现Session管理机制 (更新于2018.02.04)
- 4.1 实现了Session管理器，控制Session的添加，修改，查询，删除及其过期自动回收
- 4.2 UML图
![Session设计UML模型](https://github.com/zcv8/YM.JinLiRead/blob/master/uml/Session模型设计?raw=true)

> 参考：https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/06.0.md

### 5.重新整理项目使用Godep进行依赖管理（更新于2018.02.27）
- 5.1 在5的基础上添加Session的指定时间内过期，比如7天内免登录功能
- 5.2 使用Godep依赖管理工具管理Go项目，并调整了Go项目的文件结构
- 5.3 需要将项目签出到src/github.com/zcv8/ 目录下，然后使用 `go get github.com/tools/godep` 安装godep工具，安装成功之后就会在bin目录下出现godep的可执行文件。具体使用方式可以运行‘godep help’命令查看

### 6.使用Postgresql数据库实现对数据的处理（更新于2018.03.06）
- 6.1 使用`github.com/lib/pq`驱动连接远程云服务器数据库Postgresql

### 7.使用Delve进行Debug Go程序（更新于2018.03.07）
- 7.1 使用Delve代替log的方式，实现对程序的调试

> 参考：https://github.com/derekparker/delve?spm=a2c4e.11153959.blogcont57578.4.24c5266eNCm09t
> 中文博客：https://yq.aliyun.com/articles/57578

### 8.抽取前端展示逻辑（更新于2018.03.14）
- 8.1 整理项目为SPA程序，抽取前端业务逻辑到单独的项目
- 8.2 使用 Vue + Webpack 重构抽掉出来的前端项目

### 9.实现跨站请求处理（更新于2018.03.22）
- 9.1 前端实现过程：

（1）使用`vue-resource`进行对请求拦截，在请求之前加上`request.credentials=true;`,从而可以实现在请求的时候请求头中带Cookie。

（2）同样使用拦截器，拦截返回的请求，如果请求中存在Cookie字段，则使用`document.cookie`实现将Cookie信息存储到浏览器
- 9.2 后台实现过程：

（1）闭包实现装饰器，对每次请求的响应中，添加响应头,代码如下：
	```
		//这里不使用'*'的原因是使用11.1（1）中的请求头后，出于安全协议不允许使用'*'
		w.Header().Set("Access-Control-Allow-Origin", "http://vue.lovemoqing.com") 
		w.Header().Set("Access-Control-Allow-Headers","Cookie,Origin, X-Requested-With, Content-Type, Accept")
	```
	
 (2) 在需要返回的Cookie的请求中，将拼接好的cookie字符串返回给前端
 
 ### 10.使用Xorm重新实现数据操作逻辑（更新于2018.05.03）
 - 10.1 使用开源的第三方ORM 实现对数据的操作逻辑
 > 参考：http://www.xorm.io/docs/

### 11.使用Logrus实现对日志的统一管理（更新于2018.05.06）
- 11.1 使用开源的日志记录器Logrus，Logrus提供多样化的日志纪录
> 参考：https://github.com/sirupsen/logrus




