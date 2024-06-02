package handler

import (
	"db_proj/define"
	mstablemgrclient "db_proj/mstablemgr/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// OrderDish
// @Summary OrderDish
// @Description "点菜"
// @Tags public
// @Param table_id query string true "table_id"
// @Param dish_id formData []string true "dish_id"
// @Success 200 {json} {}
// @Router /order-dish [POST]
func HandleOrderDish(ctx *gin.Context) {
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

	dishIdStrings := ctx.PostFormArray("dish_id")
	dishIds := make([]int32, 0)

	for _, dishIdString := range dishIdStrings {
		dishId, err := strconv.ParseInt(dishIdString, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": fmt.Sprintf("Invalid dish_id, your dish_id is %s", dishIdString),
			})

			util.Log("parsing dish id error, id = %s", dishIdString)
			return
		}
		dishIds = append(dishIds, int32(dishId))
	}

	log.Println("before call mstablemgrclient.CallOrderDish", time.Now())
	resp, err := mstablemgrclient.CallOrderDish(int32(tableId), dishIds)
	log.Println("after call mstablemgrclient.CallOrderDish", time.Now())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Call mstablemgrclient.CallOrderDish error %v", err),
		})

		util.Log("Call mstablemgrclient.CallOrderDish error: %v", err)
		return
	}

	if resp.GetStatus() == define.TableIsNotInUse {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Table %d is not in use", tableId),
		})

		util.Log("Table %d is not in use", tableId)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
