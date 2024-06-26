package middleware

import (
	"db_proj/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExtractUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if len(token) == 0 {
			token = ctx.PostForm("token")
			if len(token) == 0 {
				token = ctx.GetHeader("Authorization")
			}
		}

		if len(token) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"success": "false",
				"message": "token is empty, please take token when you request server",
			})

			ctx.Abort()
			return
		}

		ok, user := util.CheckToken(token)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"success": "false",
				"message": "token invalid",
			})

			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
