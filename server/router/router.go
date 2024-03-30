package router

import (
	"github.com/gin-gonic/gin"
	"db_proj/handler"
)

func SetRouter(server *gin.Engine) {
	server.GET("/ping", handler.HandlePing)	
}