package handler

import (
	"db_proj/model"
	"db_proj/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// AddDish
// @Summary AddDish
// @Description "添加一道菜"
// @Tags public
// @Param name query string true "name"
// @Param price query number true "price"
// @Param discount query number true "discount"
// @Param detail query string false "detail"
// @Success 200 {json} {}
// @Router /admin/add-dish [GET]
func HandleAddDish(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("the name received is empty, INVALID!"),
		})
		return
	}

	detail := ctx.Query("detail")

	priceString := ctx.Query("price")
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("the price which is %s can't convert to a float(double) type, INVALID!", priceString),
		})
		return
	}

	discountString := ctx.Query("discount")
	discount, err := strconv.ParseFloat(discountString, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": fmt.Sprintf("the discount which is %s can't convert to a float(double) type, INVALID!", discountString),
		})
		return
	}

	// add
	db := model.NewMySqlConnector()
	user := model.Dish{
		Model:    gorm.Model{},
		Name:     name,
		Price:    price,
		Discount: discount,
		Detail:   detail,
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": fmt.Sprintf("create new dish into db error"),
		})

		// write into log
		util.Log("insert dish into db error, new dish = %v, error: %v", user, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
