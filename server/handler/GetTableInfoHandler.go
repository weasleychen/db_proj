package handler

import (
	mstablemgrclient "db_proj/mstablemgr/client"
	mstablemgr_service "db_proj/mstablemgr/service"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTableInfo
// @Summary GetTableInfo
// @Description "获取桌台详情，包括桌台状态、点菜等"
// @Tags public
// @Param table_id query string true "桌台id"
// @Success 200 {json} {}
// @Router /get-table-info [GET]
func HandleGetTableInfo(ctx *gin.Context) {
	tableIdString := ctx.Query("table_id")
	tableId, err := strconv.ParseInt(tableIdString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("Invalid table_id, your table_id is %s", tableIdString),
		})

		util.Log("parsing table id error, id = %s", tableIdString)
		return
	}

	resp, err := mstablemgrclient.CallGetTableInfo(int32(tableId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("Get table info failed, error: %v", err),
		})

		util.Log("get table info failed, error: %v", err)
		return
	}

	jsonString := resp.GetJsonString()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"table":   util.UnmarshalJson[mstablemgr_service.Table](jsonString),
	})
}
