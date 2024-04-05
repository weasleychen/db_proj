1.需要设定好GOROOT和GOPATH, 并且将GOPATH/bin加入环境变量PATH

# swagger api 注解模板
```go
// Ping
// @Summary ping
// @Description "ping一下服务器"
// @Tags example
// @Param message query string false "回复"
// @Success 200 {json} {}
// @Router /ping [GET]
```

# 启动脚本参数
    -s 启用swagger文档
    -p <port> 指定port为监听端口