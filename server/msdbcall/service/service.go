package service

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	"db_proj/msdbcall/proto"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type MSDBCallServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *MSDBCallServer) AddDish(ctx context.Context, req *msdbcall.AddDishReq) (*msdbcall.AddDishResp, error) {
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

func (server *MSDBCallServer) CheckUserNameUnique(ctx context.Context, req *msdbcall.CheckUserNameUniqueReq) (*msdbcall.CheckUserNameUniqueResp, error) {
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

func (server *MSDBCallServer) CheckUserPassword(ctx context.Context, req *msdbcall.CheckUserPasswordReq) (*msdbcall.CheckUserPasswordResp, error) {
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

func (server *MSDBCallServer) CreateUser(ctx context.Context, req *msdbcall.CreateUserReq) (*msdbcall.CreateUserResp, error) {
	user := model.User{
		Uin:         *req.User.Uin,
		Name:        *req.User.Name,
		Password:    *req.User.Password,
		PhoneNumber: *req.User.PhoneNumber,
		Perm:        *req.User.Perm,
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
