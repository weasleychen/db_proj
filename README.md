#### 需要在系统环境变量设定好GOROOT和GOPATH, 并且将GOPATH/bin和GOROOT/PATH加入环境变量PATH

#### 启动脚本参数
    -s 启用swagger文档
    -p <port> 指定port为监听端口

#### 安装swagger
go install github.com/go-swagger/go-swagger/cmd/swagger@latest

#### 安装go-grpc和go-grpc-gen
sudo apt install protobuf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1


