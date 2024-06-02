package service

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	model_pb "db_proj/model/proto"
	msdbcallclient "db_proj/msdbcall/client"
	"db_proj/msdbcall/proto"
	"db_proj/util"
	"encoding/json"
	"errors"
	"fmt"
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

func (server *MSDBCallServer) CheckUserPassword(ctx context.Context, req *msdbcall.CheckUserPasswordReq) (*msdbcall.CheckUserPasswordResp, error) {
	db := model.NewMySqlConnector()

	resp := msdbcall.CheckUserPasswordResp{}
	user := model.User{}

	if req.GetUin() != "" {
		if err := db.Where("uin = ?", req.GetUin()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchUin)
				return &resp, nil
			}

			// 意料之外的db错误, 直接抛给上层，以便能看到具体的错误信息
			return nil, err
		}
	} else if req.GetPhoneNumber() != "" {
		if err := db.Where("phone_number = ?", req.GetPhoneNumber()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchPhoneNumber)
				return &resp, nil
			}
			return nil, err
		}
	} else if req.GetEmail() != "" {
		if err := db.Where("email = ?", req.GetEmail()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchEmail)
				return &resp, nil
			}
			return nil, err
		}
	}

	if user.Password != *req.Password {
		resp.Status = util.NewType[int32](define.ErrorWrongPassword)
	} else {
		resp.Status = util.NewType[int32](define.OK)
	}

	return &resp, nil
}

func (server *MSDBCallServer) CreateUser(ctx context.Context, req *msdbcall.CreateUserReq) (*msdbcall.CreateUserResp, error) {
	user := model.User{
		Uin:         *req.User.Uin,
		Name:        *req.User.Name,
		Password:    *req.User.Password,
		PhoneNumber: *req.User.PhoneNumber,
		Perm:        *req.User.Perm,
		Email:       *req.User.Email,
	}

	resp := msdbcall.CreateUserResp{}

	db := model.NewMySqlConnector()
	tx := db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		resp.Code = util.NewType[int32](define.ErrorCreateUser)
		return &resp, err
	}

	tx.Commit()

	resp.Code = util.NewType[int32](define.OK)
	resp.Data = util.MarshalJsonRetPtr(user)
	return &resp, nil
}

func (server *MSDBCallServer) ModifyPassword(ctx context.Context, req *msdbcall.ModifyPasswordReq) (*msdbcall.ModifyPasswordResp, error) {
	checkUserPasswordResp, err := msdbcallclient.CallCheckUserPassword(req.GetUin(), req.GetPhoneNumber(), req.GetEmail(), req.GetOldPassword())
	if err != nil {
		return nil, err
	}

	resp := msdbcall.ModifyPasswordResp{}
	if checkUserPasswordResp.GetStatus() != define.OK {
		resp.Status = util.NewType[int32](define.ErrorWrongPassword)
		return &resp, nil
	}

	db := model.NewMySqlConnector()
	// 根据uin > phone_number > email的优先级来修改密码
	if req.GetUin() != "" {
		if err := db.Model(&model.User{}).Where("uin = ?", req.GetUin()).Update("password", req.GetNewPassword()).Error; err != nil {
			return nil, err
		}
	} else if req.GetPhoneNumber() != "" {
		if err := db.Model(&model.User{}).Where("phone_number = ?", req.GetPhoneNumber()).Update("password", req.GetNewPassword()).Error; err != nil {
			return nil, err
		}
	} else if req.GetEmail() != "" {
		if err := db.Model(&model.User{}).Where("email = ?", req.GetEmail()).Update("password", req.GetNewPassword()).Error; err != nil {
			return nil, err
		}
	}

	resp.Status = util.NewType[int32](define.OK)
	return &resp, nil
}

func (server *MSDBCallServer) GetUserInfo(ctx context.Context, req *msdbcall.GetUserInfoReq) (*msdbcall.GetUserInfoResp, error) {
	// 根据uin > phone_number > email的优先级来查询出一行数据
	db := model.NewMySqlConnector()
	user := model.User{}
	resp := msdbcall.GetUserInfoResp{}

	if req.GetUin() != "" {
		if err := db.Where("uin = ?", req.GetUin()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchUin)
				return &resp, nil
			}
			return nil, err
		}
	} else if req.GetPhoneNumber() != "" {
		if err := db.Where("phone_number = ?", req.GetPhoneNumber()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchPhoneNumber)
				return &resp, nil
			}
			return nil, err
		}
	} else if req.GetEmail() != "" {
		if err := db.Where("email = ?", req.GetEmail()).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.Status = util.NewType[int32](define.ErrorNoSuchEmail)
				return &resp, nil
			}
			return nil, err
		}
	}

	resp.Status = util.NewType[int32](define.OK)
	resp.Data = util.MarshalJsonRetPtr(user)
	return &resp, nil
}

func (server *MSDBCallServer) GetDishList(ctx context.Context, req *msdbcall.GetDishListReq) (*msdbcall.GetDishListResp, error) {
	resp := msdbcall.GetDishListResp{}

	db := model.NewMySqlConnector()
	modelDishList := make([]model.Dish, 0)
	if err := db.Find(&modelDishList).Error; err != nil {
		return nil, err
	}

	fmt.Println("size: ", len(modelDishList))

	for _, modelDish := range modelDishList {
		resp.DishList = append(resp.DishList, &model_pb.Dish{
			Id:       util.NewType[int32](int32(modelDish.ID)),
			Name:     &modelDish.Name,
			Price:    &modelDish.Price,
			Discount: &modelDish.Discount,
			Detail:   &modelDish.Detail,
		})
	}

	return &resp, nil
}

func (server *MSDBCallServer) DeleteDish(ctx context.Context, req *msdbcall.DeleteDishReq) (*msdbcall.DeleteDishResp, error) {
	db := model.NewMySqlConnector()
	resp := msdbcall.DeleteDishResp{}

	if err := db.Where("id = ?", req.GetId()).Delete(&model.Dish{}).Error; err != nil {
		return nil, err
	}

	resp.Status = util.NewType[int32](define.OK)
	return &resp, nil
}

func (server *MSDBCallServer) GetDishInfo(ctx context.Context, req *msdbcall.GetDishInfoReq) (*msdbcall.GetDishInfoResp, error) {
	resp := msdbcall.GetDishInfoResp{}

	db := model.NewMySqlConnector()

	protoDish := model_pb.Dish{}
	modelDish := model.Dish{}

	if err := db.Debug().Model(&model.Dish{}).Where("id = ?", req.GetDishId()).First(&modelDish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Status = util.NewType[int32](define.ErrorDishIdNotExist)
			return &resp, nil
		}
		return nil, err
	}

	protoDish.Id = util.NewType[int32](int32(modelDish.ID))
	protoDish.Name = &modelDish.Name
	protoDish.Price = &modelDish.Price
	protoDish.Detail = &modelDish.Detail

	resp.Dish = &protoDish
	resp.Status = util.NewType[int32](define.OK)
	return &resp, nil
}
