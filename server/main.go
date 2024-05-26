package main

import (
	"db_proj/define"
	"db_proj/handler"
	"db_proj/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

/*
	题目七：
	餐饮管理系统

	系统概述：餐饮管理系统是一套为中小型餐饮企业提供的用于点餐和结算的数据库管理系统。

	基本要求：
	（1）桌台管理
	（2）就餐管理：包括开台、点餐、结账等
	（3）菜单管理 TODO
	（4）用户管理 completed
	（5）营业额管理
	（6）系统维护
	（7）可根据上述需求进行合适的调整和改变
*/

/* 	待选:
- 异步日志
*/

/*
	1. 模块内部不打日志，日志只在gin层打
*/

func init() {
	os.Mkdir(fmt.Sprintf("%s/server/log", define.ProjectPath), os.ModePerm)
}

func StartServer(port int) *gin.Engine {
	server := gin.Default()

	router.SetRouter(server)

	err := server.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("gin.Engine::Run() Error")
		return nil
	}

	return server
}

func main() {
	handler.ParseCommandLine()

	if define.DebugMode {
		fmt.Printf("port: %d\nUseSwagger:%v\n", define.Port, define.UseSwagger)
	}

	StartServer(define.Port)
}
