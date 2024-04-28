package service

import (
	"context"
	"db_proj/model"
	msdbcall "db_proj/msdbcall/proto"
	"errors"
	"gorm.io/gorm"
)

type CheckUserPasswordServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *CheckUserPasswordServer) CheckUserPassword(ctx context.Context, req *msdbcall.CheckUserPasswordReq) (*msdbcall.CheckUserPasswordResp, error) {
	db := model.NewMySqlConnector()

	var success bool
	resp := msdbcall.CheckUserPasswordResp{Success: &success}

	user := model.User{}
	if err := db.Where("name = ? and password = ?", *req.Name, *req.Password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			*resp.Success = false
			return &resp, nil
		}

		return nil, err
	}

	*resp.Success = true
	return &resp, nil
}
