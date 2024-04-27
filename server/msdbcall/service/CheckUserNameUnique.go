package service

import (
	"context"
	"db_proj/model"
	msdbcall "db_proj/msdbcall/proto"
	"errors"
	"gorm.io/gorm"
)

type CheckUserNameUniqueServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *CheckUserNameUniqueServer) CheckUserNameUnique(ctx context.Context, req *msdbcall.CheckUserNameUniqueReq) (*msdbcall.CheckUserNameUniqueResp, error) {
	db := model.NewMySqlConnector()

	var unique bool
	resp := msdbcall.CheckUserNameUniqueResp{}
	resp.Unique = &unique

	user := model.User{}
	if err := db.Where(&model.User{Name: *req.Name}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			*resp.Unique = false
			return &resp, nil
		}
		return nil, err
	}

	*resp.Unique = true
	return &resp, nil
}
