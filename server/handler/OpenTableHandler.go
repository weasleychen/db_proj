package handler

import (
	"db_proj/define"
	mstablemgrclient "db_proj/mstablemgr/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// OpenTable
// @Summary OpenTable
// @Description "开台"
// @Tags public
// @Param table_id query string true "table_id"
// @Success 200 {json} {}
// @Router /admin/open-table [GET]
func HandleOpenTable(ctx *gin.Context) {
	tableIdString := ctx.Query("table_id")

	tableId, err := strconv.ParseInt(tableIdString, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("tableId err: %v", err),
		})

		util.Log("Error tableId: %v", tableIdString)
		return
	}

	resp, err := mstablemgrclient.CallOpenTable(int32(tableId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("CallOpenTable err: %v", err),
		})

		util.Log("Error CallOpenTable: %v", err)
		return
	}

	if *resp.Status == define.ErrorTableIsOpened {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Error, TableIsOpened, tableId: %v", tableId),
		})

		util.Log("Error, TableIsOpened, tableId: %v", tableId)
		return
	}

	if *resp.Status == define.ErrorTableIdNotExist {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Error ErrorTableIdNotExist, tableId: %v", tableId),
		})

		util.Log("Error ErrorTableIdNotExist, tableId: %v", tableId)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})

	util.Log("OpenTable success, tableId: %v", tableId)
}
