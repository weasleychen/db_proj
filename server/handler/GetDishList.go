package handler

import (
	msdbcallclient "db_proj/msdbcall/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDishList
// @Summary GetDishList
// @Description "获取菜品列表"
// @Tags public
// @Success 200 {json} {}
// @Router /get-dish-list [GET]
func HandleGetDishList(ctx *gin.Context) {
	resp, err := msdbcallclient.CallGetDishList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("error: %v", err),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  "true",
		"dishList": resp.GetDishList(),
	})

}
