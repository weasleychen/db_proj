package handler

import (
	mstablemgrclient "db_proj/mstablemgr/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTablesStatus
// @Summary GetTablesStatus
// @Description "获得全部桌台状态"
// @Tags public
// @Success 200 {json} {}
// @Router /admin/get-tables-status [GET]
func HandleGetTablesStatus(ctx *gin.Context) {
	resp, err := mstablemgrclient.CallGetTablesStatus()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("call mstablemgr.GetTablesStatus failed, err: %v", err),
		})

		util.Log("call mstablemgr.GetTablesStatus failed, err: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"tables":  resp.TableList,
	})

	util.Log("call mstablemgr.GetTablesStatus success, resp: %v", resp)
}
