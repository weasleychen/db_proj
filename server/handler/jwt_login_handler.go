package handler

import (
	"db_proj/define"
	"db_proj/model"
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
func HandlerJWTLogin(ctx *gin.Context) {
	tokenString := ctx.PostForm("token")

	user, err := model.TokenRedisHandler.GetRedisClient().Get(define.DefaultRedisContext, tokenString).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
			"user":    user,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
		})
	}
}
