package handler

import "github.com/gin-gonic/gin"

func HandlePing(ctx *gin.Context) {
	ctx.JSON(200, gin.H {
		"success": "true",
	})
}