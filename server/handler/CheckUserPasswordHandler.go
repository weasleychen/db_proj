package handler

import (
	"db_proj/define"
	msdbcallclient "db_proj/msdbcall/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CallCheckUserPassword
// @Summary CallCheckUserPassword
// @Description "检查md5加密用户密码是否正确"
// @Tags public
// @Param uin query string false "uin"
// @Param phone_number query string false "phone_number"
// @Param email query string false "email"
// @Param password formData string true "password"
// @Success 200 {json} {}
// @Router /check-user-password [GET]
func HandleCheckUserPassword(ctx *gin.Context) {
	if ctx.Query("uin") == "" && ctx.Query("phone_number") == "" && ctx.Query("email") == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "argument invalid, you must choose one of (uin, phone_number, email)",
		})
	}

	resp, err := msdbcallclient.CallCheckUserPassword(
		ctx.Query("uin"),
		ctx.Query("phone_number"),
		ctx.Query("email"),
		ctx.PostForm("password"),
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("error: %v", err),
		})
		return
	}

	if *resp.Status == define.ErrorNoSuchUin {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "no such uin",
		})
		return
	}

	if *resp.Status == define.ErrorNoSuchPhoneNumber {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "no such phone number",
		})
		return
	}

	if *resp.Status == define.ErrorNoSuchEmail {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "no such email",
		})
		return
	}

	if *resp.Status == define.ErrorWrongPassword {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "password incorrect",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
