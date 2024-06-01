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
	server.Use(middleware.Cors())

	//router for swagger
	if define.UseSwagger {
		server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	admin := server.Group("/admin", middleware.CheckAdmin())

	server.GET("/ping", handler.HandlePing)

	server.POST("/register", handler.HandleRegister)
	server.POST("/login", handler.HandleLogin)
	server.POST("/login-by-token", handler.HandleJWTLogin)
	server.GET("/get-tables-status", handler.HandleGetTablesStatus)
	server.POST("/modify-password", handler.HandleModifyPassword)
	server.GET("/get-dish-list", handler.HandleGetDishList)

	admin.GET("/add-dish", handler.HandleAddDish)
	admin.GET("/open-table", handler.HandleOpenTable)
	admin.GET("/complete-table", handler.HandleCompleteTable)
	admin.GET("/add-table", handler.HandleAddTable)
	admin.GET("/del-table", handler.HandleDelTable)
	admin.GET("/delete-dish", handler.HandleDeleteDish)
}
