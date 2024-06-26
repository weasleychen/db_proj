package handler

import (
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
// @Summary Register
// @Description "注册新用户"
// @Tags public
// @Param name formData string true "用户名"
// @Param password formData string true "MD5加密密码"
// @Param phone_number formData string true "手机"
// @Param email formData string true "邮箱"
// @Success 200 {json} {}
// @Router /register [POST]
func HandleRegister(ctx *gin.Context) {
	user := model.User{
		Uin:         util.GenNewUin(),
		Name:        ctx.PostForm("name"),
		Password:    ctx.PostForm("password"),
		PhoneNumber: ctx.PostForm("phone_number"),
		Email:       ctx.PostForm("email"),
		Perm:        define.NormalPerm,
	}

	if user.PhoneNumber == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty phone number",
		})

		util.Log("empty phone number")
		return
	}

	if user.Email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty email",
		})

		util.Log("empty email")
		return
	}

	if user.Password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty password",
		})

		util.Log("empty password")
		return
	}

	resp, err := msdbcallclient.CallCreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("err: %v", err),
		})

		util.Log("create new user error, new user = %v, err: %v, code = %d", user, err)
		return
	}

	if resp.GetCode() == define.ErrorDuplicatePhoneNumber {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "duplicate phone_number or email",
		})

		util.Log("duplicate phone_number or email")
		return
	}

	json.Unmarshal([]byte(*resp.Data), &user)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"user":    user,
		"token":   util.GenJWT(util.MarshalJson(&user)),
	})
}
