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

func SetRouter(user *gin.Engine) {
	user.Use(middleware.Cors())

	//router for swagger
	if define.UseSwagger {
		user.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	admin := user.Group("/admin", middleware.ExtractUser(), middleware.CheckAdmin())

	user.GET("/ping", handler.HandlePing)

	// normal user without token
	user.GET("/get-dish-list", handler.HandleGetDishList)
	user.POST("/register", handler.HandleRegister)
	user.POST("/login", handler.HandleLogin)
	user.POST("/login-by-token", handler.HandleJWTLogin)

	// normal user with token
	tokenUser := user.Group("/", middleware.ExtractUser())
	tokenUser.GET("/get-table-info", handler.HandleGetTableInfo)
	tokenUser.GET("/get-dish-info", handler.HandleGetDishInfo)
	tokenUser.GET("/open-table", handler.HandleOpenTable)
	tokenUser.GET("/complete-table", handler.HandleCompleteTable)
	tokenUser.POST("/modify-password", handler.HandleModifyPassword)
	tokenUser.POST("/order-dish", handler.HandleOrderDish)

	// admin must with token
	admin.GET("/get-tables-status", handler.HandleGetTablesStatus)
	admin.GET("/add-dish", handler.HandleAddDish)
	admin.GET("/add-table", handler.HandleAddTable)
	admin.GET("/del-table", handler.HandleDelTable)
	admin.GET("/delete-dish", handler.HandleDeleteDish)
	admin.GET("/get-consume-record", handler.HandleGetConsumeRecord)
	admin.POST("/admin-add-user", handler.HandleAdminAddUser)

	user.Static("/static", "/home/stdforces/static")
}
