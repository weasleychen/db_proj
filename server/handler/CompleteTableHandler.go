package handler

import (
	"db_proj/define"
	"db_proj/model"
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
// @Param token query string true "token"
// @Success 200 {json} {}
// @Router /complete-table [GET]
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

	userRaw, _ := ctx.Get("user")
	user, ok := userRaw.(model.User)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("unexpected type assertion"),
		})

		util.Log("unexpected type assertion")
		return
	}

	resp, err := mstablemgrclient.CallCompleteTable(user.Uin, int32(tableId))
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
