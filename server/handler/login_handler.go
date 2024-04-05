package handler

import (
	"db_proj/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	model.NewMySqlConnector().Where(&user, "name").Find(&user)

	if user.Password == ctx.PostForm("password") {
		ctx.JSON(200, gin.H{
			"success": "true",
		})
	} else {
		ctx.JSON(200, gin.H{
			"success": "false",
		})
	}
}
