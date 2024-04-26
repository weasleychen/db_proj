package service

import (
	"context"
	"db_proj/model"
	msdbcall "db_proj/msdbcall/proto"
	"encoding/json"
)

type CreateUserServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *CreateUserServer) CreateUser(ctx context.Context, req *msdbcall.CreateUserReq) (*msdbcall.CreateUserResp, error) {
	user := model.User{
		Uin:         *req.Uin,
		Name:        *req.Name,
		Password:    *req.Password,
		PhoneNumber: *req.PhoneNumber,
		Perm:        *req.Perm,
	}

	db := model.NewMySqlConnector()
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	jsonString := string(jsonBytes)

	resp := msdbcall.CreateUserResp{}
	resp.Data = &jsonString

	return &resp, nil
}
