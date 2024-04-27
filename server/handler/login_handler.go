package handler

import (
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// LoginByName
// @Summary LoginByName
// @Description "登录"
// @Tags public
// @Param name formData string true "用户名"
// @Param password formData string true "MD5加密密码"
// @Success 200 {json} {}
// @Router /login-by-name [POST]
func HandleLoginByName(ctx *gin.Context) {
	user := model.User{
		Model: gorm.Model{},
		Name:  ctx.PostForm("name"),
		Perm:  0,
	}

	resp, err := msdbcallclient.CheckUserPassword(user.Name, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
		})

		util.Log("CheckUserPassword Error, err: %v", err)
		return
	}

	if *resp.Success {
		jsonBytes, err := json.Marshal(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		jsonString := string(jsonBytes)
		tokenString := util.GenJWT(jsonString)

		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
			"token":   tokenString,
			"user":    user,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"ret":     define.ErrorWrongPassword,
		})
	}
}
