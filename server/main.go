package main

import (
	"db_proj/cmdlinehandler"
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
	port := cmdlinehandler.ParseCommandLine()
	StartServer(port)
	
}
