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

// AddTable
// @Summary AddTable
// @Description "加台"
// @Tags public
// @Param table_id formData string true "table_id"
// @Success 200 {json} {}
// @Router /add-table [GET]
func HandleAddTable(ctx *gin.Context) {
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

	resp, err := mstablemgrclient.CallAddTable(int32(tableId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Call mstablemgr.AddTable error"),
		})

		util.Log("Call mstablemgr.AddTable error")
		return
	}

	if *resp.Status == define.ErrorTableIdExist {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Error Table already exist, your table_id: %s", tableIdString),
		})

		util.Log("Error Table already exist, your table_id: %s", tableIdString)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
