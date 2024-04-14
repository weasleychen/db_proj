package router

import (
	"db_proj/define"
	"db_proj/handler"

	_ "db_proj/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter(server *gin.Engine) {
	//router for swagger
	if define.UseSwagger {
		server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	server.GET("/ping", handler.HandlePing)

	server.POST("/register", handler.HandleRegister)
	server.POST("/login-by-name", handler.HandleLoginByName)
	server.POST("/login-by-token", handler.HandlerJWTLogin)
}
