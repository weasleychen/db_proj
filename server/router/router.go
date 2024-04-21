package router

import (
	"db_proj/define"
	"db_proj/handler"
	"db_proj/middleware"

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

	admin := server.Group("/admin", middleware.CheckAdmin())

	server.GET("/ping", handler.HandlePing)

	server.POST("/register", handler.HandleRegister)
	server.POST("/login-by-name", handler.HandleLoginByName)
	server.POST("/login-by-token", handler.HandleJWTLogin)

	admin.GET("/add-dish", handler.HandleAddDish)
}
