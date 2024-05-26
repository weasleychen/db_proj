package main

import (
	"db_proj/define"
	"db_proj/msdbcall/proto"
	"db_proj/msdbcall/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterNewHandler(address string, handler msdbcall.MSDBCallServer) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("listen on ", address, " error")
		return err
	}

	server := grpc.NewServer()
	msdbcall.RegisterMSDBCallServer(server, handler)

	return server.Serve(listener)
}

func main() {
	RegisterNewHandler(define.MSDBCallIp+":"+define.MSDBCallPort, &service.MSDBCallServer{})
}
