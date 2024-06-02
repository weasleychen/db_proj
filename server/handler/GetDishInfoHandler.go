package handler

import (
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetDishInfo
// @Summary GetDishInfo
// @Description "获得菜品详情"
// @Tags public
// @Param dish_id query string true "dish_id"
// @Success 200 {json} {}
// @Router /get-dish-info [GET]
func HandleGetDishInfo(ctx *gin.Context) {
	dishIdString := ctx.Query("dish_id")
	dishId, err := strconv.Atoi(dishIdString)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("dish_id is invalid, error: %v", err),
		})

		util.Log("dish_id is invalid, error: %v", err)
		return
	}

	resp, err := msdbcallclient.CallGetDishInfo(int32(dishId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("CallGetDishInfo error: %v", err),
		})

		util.Log("CallGetDishInfo error: %v", err)
		return
	}

	if resp.GetStatus() == define.ErrorDishIdNotExist {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("dish_id %d not exist", dishId),
		})

		util.Log("dish_id %d not exist", dishId)
		return
	}

	modelDish := model.Dish{}
	modelDish.ID = uint(resp.Dish.GetId())
	modelDish.Name = resp.Dish.GetName()
	modelDish.Price = resp.Dish.GetPrice()
	modelDish.Discount = resp.Dish.GetDiscount()
	modelDish.Detail = resp.Dish.GetDetail()

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"dish":    modelDish,
	})
}
