package handler

import (
	"db_proj/define"
	"db_proj/model"
	"db_proj/util"
	"github.com/gin-gonic/gin"
	"log"
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

	log.Println(user)

	conn := model.NewMySqlConnector()
	conn.Create(&user)
	if conn.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
		})
	}
}
