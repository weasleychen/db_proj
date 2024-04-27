package handler

import (
	"db_proj/define"
	"db_proj/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginByJWT
// @Summary LoginByJWT
// @Description "登录"
// @Tags public
// @Param token formData string true "token"
// @Success 200 {json} {}
// @Router /login-by-token [POST]
func HandleJWTLogin(ctx *gin.Context) {
	tokenString := ctx.PostForm("token")

	valid, user := util.CheckToken(tokenString)
	if valid {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
			"user":    user,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"ret":     define.ErrorTokenInvalid,
		})
	}
}
