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
	"strconv"
)

// AdminAddUser
// @Summary AdminAddUser
// @Description "管理员新增用户"
// @Tags public
// @Param name formData string true "用户名"
// @Param phone_number formData string true "手机号"
// @Param email formData string true "邮箱"
// @Param password formData string true "MD5加密密码"
// @Param perm formData string true "权限"
// @Success 200 {json} {}
// @Router /admin/admin-add-user [POST]
func HandleAdminAddUser(ctx *gin.Context) {
	permString := ctx.PostForm("perm")
	if permString == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty perm",
		})

		util.Log("empty perm")
		return
	}

	perm, err := strconv.ParseInt(permString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Parse perm error, your perm: %s", permString),
		})

		util.Log("Parse perm error, your perm: %d", permString)
		return
	}

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phoneNumber := ctx.PostForm("phone_number")
	email := ctx.PostForm("email")

	if name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty name",
		})

		util.Log("empty name")
		return
	}

	if password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty password",
		})

		util.Log("empty password")
		return
	}

	if phoneNumber == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty phone_number",
		})

		util.Log("empty phone_number")
		return
	}

	if email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "empty email",
		})

		util.Log("empty email")
		return
	}

	user := model.User{
		Uin:         util.GenNewUin(),
		Name:        name,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
		Perm:        int32(perm),
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
			"message": "duplicate phone_number",
		})

		util.Log("duplicate phone_number: %v", err)
		return
	}

	json.Unmarshal([]byte(*resp.Data), &user)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"user":    user,
		"token":   util.GenJWT(util.MarshalJson(&user)),
	})
}
