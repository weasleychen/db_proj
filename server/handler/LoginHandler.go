package handler

import (
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary Login
// @Description "登录"
// @Tags public
// @Param phone_number query string true "手机号"
// @Param password formData string true "MD5加密密码"
// @Success 200 {json} {}
// @Router /login [POST]
func HandleLogin(ctx *gin.Context) {
	// 支持三种登录方式，不过swagger只开放了phone_number
	resp, err := msdbcallclient.CallCheckUserPassword(ctx.Query("uin"), ctx.Query("phone_number"), ctx.Query("email"), ctx.PostForm("password"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
		})

		util.Log("CallCheckUserPassword Error, err: %v", err)
		return
	}

	if resp.GetStatus() == define.OK {
		getUserInfoResp, err := msdbcallclient.CallGetUserInfo(ctx.Query("uin"), ctx.Query("phone_number"), ctx.Query("email"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
			})

			util.Log("CallGetUserInfo Error, err: %v", err)
			return
		}

		jsonString := getUserInfoResp.GetData()
		tokenString := util.GenJWT(jsonString)

		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
			"token":   tokenString,
			"user":    util.UnmarshalJsonRetPrt[model.User](jsonString),
		})
	} else {
		util.Log("check password error, status: %v", resp.GetStatus())
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"status":  define.ErrorWrongPassword,
		})
	}
}
