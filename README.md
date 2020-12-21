## myBlog 服务端
### 安装服务端
`go get github.com/SYSU-myBlog/Server`

如果出现以下报错信息：
```bash
go get github.com/SYSU-myBlog/Server
go: downloading github.com/SYSU-myBlog/Server v0.0.0-20201220135209-8db1989bc9ba
go: github.com/SYSU-myBlog/Server upgrade => v0.0.0-20201220135209-8db1989bc9ba
go get: github.com/SYSU-myBlog/Server@v0.0.0-20201220135209-8db1989bc9ba requires
	github.com/SYSU-myBlog/Server/App@v0.0.0: reading https://goproxy.io/github.com/%21s%21y%21s%21u-my%21blog/%21server/%21app/@v/v0.0.0.mod: 404 Not Found
	server response: not found: unknown revision App0.0.0
```
可能是你之前设置过代理，运行`go env -w GO111MODULE=off`，然后再运行go get命令即可。

如果仍然不行，则直接clone仓库到本地。

### 安装和配置mongodb数据库
ubuntu mongodb安装和使用
[https://www.cnblogs.com/weihu/p/8570083.html](https://www.cnblogs.com/weihu/p/8570083.html)

使用默认配置即可

然后创建一个新的数据库，命名为myblog，在该数据库下新建四个表格，分别为user、article、comment、like。

如果想要有自己的数据库设置，可以查看以下main.go文件，在其中相应地作修改即可。

### 运行
进入Server文件夹，`go run main.go`



