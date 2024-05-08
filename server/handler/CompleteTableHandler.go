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

// CompleteTable
// @Summary CompleteTable
// @Description "结台"
// @Tags public
// @Param table_id query string true "table_id"
// @Success 200 {json} {}
// @Router /admin/complete-table [GET]
func HandleCompleteTable(ctx *gin.Context) {
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

	resp, err := mstablemgrclient.CallCompleteTable(int32(tableId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("error: %v", err),
		})

		util.Log("call mstablemgr error, err: %v", err)
		return
	}

	if *resp.Status == define.ErrorTableIdNotExist {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "ErrorTableIdNotExist",
		})

		util.Log("Error ErrorTableIdNotExist, tableId: %v", tableId)
		return
	}

	if *resp.Status == define.ErrorTableIsClosed {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("ErrorTableIsClosed, your table_id: %s", tableIdString),
		})

		util.Log("Error ErrorTableIsClosed, tableId: %v", tableId)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})

	util.Log("complete table success, tableId: %v", tableId)
}