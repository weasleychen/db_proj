package service

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	msdbcall "db_proj/msdbcall/proto"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
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

	resp := msdbcall.CreateUserResp{}
	code := int32(define.OK)
	resp.Code = &code

	db := model.NewMySqlConnector()
	tx := db.Begin()

	if err := tx.Where("name = ?", user.Name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Create(&user).Error; err != nil {
				return nil, err
			}

			tx.Commit()

			jsonBytes, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}
			jsonString := string(jsonBytes)

			resp.Data = &jsonString

			return &resp, nil
		}
		return nil, err
	}

	*resp.Code = define.ErrorDuplicateUserName
	return &resp, nil
}
