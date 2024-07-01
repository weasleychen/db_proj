package msdbcallclient

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	proto "db_proj/model/proto"
	"db_proj/msdbcall/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func GetMSDBCallClient() (func(), *msdbcall.MSDBCallClient, *context.Context) {
	conn, err := grpc.Dial(
		define.MSDBCallIp+":"+define.MSDBCallPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial err")
	}

	msdbcallclient := msdbcall.NewMSDBCallClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	callback := func() {
		conn.Close()
		cancel()
	}

	return callback, &msdbcallclient, &ctx
}

func CallCreateUser(user model.User) (*msdbcall.CreateUserResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.CreateUserReq{}
	req.User = new(proto.User)

	req.User.Uin = &user.Uin
	req.User.Name = &user.Name
	req.User.Password = &user.Password
	req.User.PhoneNumber = &user.PhoneNumber
	req.User.Perm = &user.Perm
	req.User.Email = &user.Email

	return (*client).CreateUser(*ctx, &req)
}

func CallAddDish(dish model.Dish) (*msdbcall.AddDishResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.AddDishReq{}
	req.Dish = new(proto.Dish)

	req.Dish.Name = &dish.Name
	req.Dish.Price = &dish.Price
	req.Dish.Discount = &dish.Discount
	req.Dish.Detail = &dish.Detail

	return (*client).AddDish(*ctx, &req)
}

func CallCheckUserPassword(uin, phoneNumber, email, password string) (*msdbcall.CheckUserPasswordResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.CheckUserPasswordReq{}
	req.Uin = &uin
	req.PhoneNumber = &phoneNumber
	req.Email = &email
	req.Password = &password

	return (*client).CheckUserPassword(*ctx, &req)
}

func CallModifyPassword(uin, phoneNumber, email, oldPassword, newPassword string) (*msdbcall.ModifyPasswordResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.ModifyPasswordReq{}
	req.Uin = &uin
	req.PhoneNumber = &phoneNumber
	req.Email = &email
	req.OldPassword = &oldPassword
	req.NewPassword = &newPassword

	return (*client).ModifyPassword(*ctx, &req)
}

func CallGetUserInfo(uin, phoneNumber, email string) (*msdbcall.GetUserInfoResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.GetUserInfoReq{}
	req.Uin = &uin
	req.PhoneNumber = &phoneNumber
	req.Email = &email

	return (*client).GetUserInfo(*ctx, &req)
}

func CallGetDishList() (*msdbcall.GetDishListResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	return (*client).GetDishList(*ctx, &msdbcall.GetDishListReq{})
}

func CallDeleteDish(dishId int32) (*msdbcall.DeleteDishResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	req := msdbcall.DeleteDishReq{}
	req.Id = &dishId

	return (*client).DeleteDish(*ctx, &req)
}

func CallGetDishInfo(dishId int32) (*msdbcall.GetDishInfoResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	return (*client).GetDishInfo(*ctx, &msdbcall.GetDishInfoReq{DishId: &dishId})
}

func CallGetUserDiscount(uin string) (*msdbcall.GetUserDiscountResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	return (*client).GetUserDiscount(*ctx, &msdbcall.GetUserDiscountReq{Uin: &uin})
}

func CallStoreConsumeRecord(data string) (*msdbcall.StoreConsumeRecordResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	return (*client).StoreConsumeRecord(*ctx, &msdbcall.StoreConsumeRecordReq{Data: &data})
}

func CallGetConsumeRecord(start, end uint32) (*msdbcall.GetConsumeRecordResp, error) {
	callback, client, ctx := GetMSDBCallClient()
	defer callback()

	return (*client).GetConsumeRecord(*ctx, &msdbcall.GetConsumeRecordReq{Start: &start, End: &end})
}
