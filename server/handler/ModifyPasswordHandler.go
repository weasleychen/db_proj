package handler

import (
	"db_proj/define"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ModifyPassword
// @Summary ModifyPassword
// @Description "修改密码"
// @Tags public
// @Param uin query string false "uin"
// @Param phone_number query string false "phone_numer"
// @Param email query string false "email"
// @Param old_password formData string true "old_password"
// @Param new_password formData string true "new_password"
// @Success 200 {json} {}
// @Router /modify-password [POST]
func HandleModifyPassword(ctx *gin.Context) {
	uin := ctx.PostForm("uin")
	phoneNumber := ctx.PostForm("phone_number")
	email := ctx.PostForm("email")

	if uin == "" && phoneNumber == "" && email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "arguments invalid, you must choose one of (uin, phone_number, email)",
		})

		util.Log("arguments invalid, you must choose one of (uin, phone_number, email)")
		return
	}

	resp, err := msdbcallclient.CallModifyPassword(uin, phoneNumber, email, ctx.PostForm("old_password"), ctx.PostForm("new_password"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("error: %v", err),
		})

		util.Log("error: %v", err)
		return
	}

	if resp.GetStatus() == define.ErrorWrongPassword {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "oldPassword wrong",
		})

		util.Log("oldPassword wrong")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
