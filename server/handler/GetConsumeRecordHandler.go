package handler

import (
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTurnover
// @Summary GetTurnover
// @Description "获取营业额"
// @Tags public
// @Param start query string true "start"
// @Param end query string true "end"
// @Success 200 {json} {}
// @Router /admin/get-consume-record [GET]
func HandleGetConsumeRecord(ctx *gin.Context) {
	startString, endString := ctx.Query("start"), ctx.Query("end")

	start, err := strconv.ParseUint(startString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("error parse start, your start: %s", startString),
		})

		return
	}

	end, err := strconv.ParseUint(endString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("error parse end, your end: %s", endString),
		})

		return
	}

	resp, err := msdbcallclient.CallGetTurnover(uint32(start), uint32(end))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	array := make([]model.ConsumeRecordJson, 0)
	for _, value := range resp.GetData() {
		array = append(array, util.UnmarshalJson[model.ConsumeRecordJson](value))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    array,
	})
}
