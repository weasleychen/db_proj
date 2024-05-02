package service

import (
	"context"
	"db_proj/model"
	"db_proj/msdbcall/proto"
	"encoding/json"
)

type AddDishServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *AddDishServer) AddDish(ctx context.Context, req *msdbcall.AddDishReq) (*msdbcall.AddDishResp, error) {
	dish := model.Dish{}
	dish.Name = *req.Dish.Name
	dish.Price = *req.Dish.Price
	dish.Discount = *req.Dish.Discount
	dish.Detail = *req.Dish.Detail

	db := model.NewMySqlConnector()
	if err := db.Create(&dish).Error; err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(dish)
	if err != nil {
		return nil, err
	}
	jsonString := string(jsonBytes)

	resp := msdbcall.AddDishResp{}
	resp.Data = &jsonString

	return &resp, nil
}
