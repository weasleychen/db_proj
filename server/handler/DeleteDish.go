package handler

import (
	"db_proj/define"
	msdbcallclient "db_proj/msdbcall/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteDish
// @Summary DeleteDish
// @Description "删除一道菜，如果dish_id不存在，不会返回失败"
// @Tags public
// @Param dish_id query string true "dish_id"
// @Success 200 {json} {}
// @Router /admin/delete-dish [GET]
func HandleDeleteDish(ctx *gin.Context) {
	dishIdString := ctx.Query("dish_id")
	dishId, err := strconv.ParseInt(dishIdString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "fasle",
			"message": fmt.Sprintf("dish_id is invalid, error: %v", err),
		})
		return
	}

	resp, err := msdbcallclient.CallDeleteDish(int32(dishId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("delete dish failed, error: %v", err),
		})
		return
	}

	if resp.GetStatus() == define.OK {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "true",
		})
	}
}
