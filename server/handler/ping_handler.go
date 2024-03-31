package handler

import "github.com/gin-gonic/gin"

// Ping
// @Summary ping
// @Description "ping一下服务器"
// @Tags example
// @Param message query string false "回复"
// @Success 200 {json} {}
// @Router /ping [GET]
func HandlePing(ctx *gin.Context) {
	message, _ := ctx.GetQuery("message")
	ctx.JSON(200, gin.H{
		"success": "true",
		"data":    "HelloWorld",
		"message": message,
	})
}
