## myBlog 服务端
### 安装
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

### 运行
进入Server文件夹，`go run main.go`


