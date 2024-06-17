package middleware

import (
	"db_proj/define"
	"db_proj/model"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRaw, _ := ctx.Get("user")
		user, ok := userRaw.(model.User)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": fmt.Sprintf("unexpected type assertion"),
			})

			util.Log("unexpected type assertion")
			return
		}

		if user.Perm != define.AdminPerm {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": fmt.Sprintf("permission denied"),
			})

			util.Log("permission denied, user: %v", user)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
