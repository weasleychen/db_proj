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

// DelTable
// @Summary DelTable
// @Description "减台"
// @Tags public
// @Param table_id query string true "table_id"
// @Success 200 {json} {}
// @Router /admin/del-table [GET]
func HandleDelTable(ctx *gin.Context) {
	tableIdString := ctx.Query("table_id")

	tableId, err := strconv.ParseInt(tableIdString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("table_id error, receive table_id: %s", tableIdString),
		})

		util.Log("table_id error, receive table_id: %s", tableIdString)
		return
	}

	resp, err := mstablemgrclient.CallDelTable(int32(tableId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("CallDelTable err: %v", err),
		})

		util.Log("Error CallDelTable: %v", err)
		return
	}

	if *resp.Status == define.ErrorTableIdNotExist {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Error, Table not exist, your table_id: %s", tableIdString),
		})

		util.Log("Error, Table not exist, your table_id: %s", tableIdString)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
