package main

import (
	"db_proj/cmdlinehandler"
	"db_proj/define"
	"db_proj/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(port int) *gin.Engine {
	server := gin.Default()

	router.SetRouter(server)

	server.Run(fmt.Sprintf("localhost:%d", port))
	return server
}

func main() {
	cmdlinehandler.ParseCommandLine()

	if define.DebugMode {
		fmt.Printf("port: %d\nUseSwagger:%v\n", define.Port, define.UseSwagger)
	}

	StartServer(define.Port)
}
