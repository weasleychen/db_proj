package handler

import (
	"db_proj/define"
	msdbcallclient "db_proj/msdbcall/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckUserPassword
// @Summary CheckUserPassword
// @Description "检查md5加密用户密码是否正确"
// @Tags public
// @Param uin formData string false "uin"
// @Param phone_number formData string false "phone_number"
// @Param email formData string false "email"
// @Param password formData string true "password"
// @Success 200 {json} {}
// @Router /check-user-password [POST]
func HandleCheckUserPassword(ctx *gin.Context) {
	uin := ctx.PostForm("uin")
	phoneNumber := ctx.PostForm("phone_number")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if uin == "" && phoneNumber == "" && email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "argument invalid, you must choose one of (uin, phone_number, email)",
		})

		return
	}

	resp, err := msdbcallclient.CallCheckUserPassword(uin, phoneNumber, email, password)

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

	if resp.GetStatus() == define.ErrorArgsInvalid {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "invalid args",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
