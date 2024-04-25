package service

import (
	"context"
	"db_proj/msdbcall/proto"
)

type AddNewDishServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *AddNewDishServer) AddNewDish(ctx context.Context, req *msdbcall.AddNewDishReq) (*msdbcall.AddNewDishResp, error) {
	// TODO
	return nil, nil
}
