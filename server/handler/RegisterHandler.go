package handler

import (
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
// @Summary Register
// @Description "注册新用户"
// @Tags public
// @Param name formData string true "用户名"
// @Param password formData string true "MD5加密密码"
// @Param phone_number formData string false "手机"
// @Success 200 {json} {}
// @Router /register [POST]
func HandleRegister(ctx *gin.Context) {
	user := model.User{
		Uin:         util.GenNewUin(),
		Name:        ctx.PostForm("name"),
		Password:    ctx.PostForm("password"),
		PhoneNumber: ctx.PostForm("phone_number"),
		Perm:        define.NormalPerm,
	}

	resp, err := msdbcallclient.CallCreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
		})

		util.Log("create new user error, new user = %v, err: %v, code = %d", user, err)
		return
	}

	if *resp.Code == define.ErrorCreateUser {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": "ErrorCreateUser, 联系后端问一下具体是什么错误",
		})

		return
	}

	json.Unmarshal([]byte(*resp.Data), &user)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"user":    user,
	})
}
