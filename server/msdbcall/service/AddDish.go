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
	dish.Name = *req.Name
	dish.Price = *req.Price
	dish.Discount = *req.Discount
	dish.Detail = *req.Detail

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
