package client

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	"db_proj/msdbcall/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func GetMMDemoClient(address string) (func(), *msdbcall.MSDBCallClient, *context.Context) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err")
	}

	mmdemoclient := msdbcall.NewMSDBCallClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	callback := func() {
		conn.Close()
		cancel()
	}

	return callback, &mmdemoclient, &ctx
}

func CallCreateUser(user model.User) (*msdbcall.CreateUserResp, error) {
	callback, client, ctx := GetMMDemoClient(define.MSDBCallIp + ":" + define.MSDBCallCreateUserPort)
	defer callback()

	req := msdbcall.CreateUserReq{}
	req.User.Uin = &user.Uin
	req.User.Name = &user.Name
	req.User.Password = &user.Password
	req.User.PhoneNumber = &user.PhoneNumber
	req.User.Perm = &user.Perm

	return (*client).CreateUser(*ctx, &req)
}

func CallAddDish(dish model.Dish) (*msdbcall.AddDishResp, error) {
	callback, client, ctx := GetMMDemoClient(define.MSDBCallIp + ":" + define.MSDBCallAddDishPort)
	defer callback()

	req := msdbcall.AddDishReq{}
	req.Dish.Name = &dish.Name
	req.Dish.Price = &dish.Price
	req.Dish.Discount = &dish.Discount
	req.Dish.Detail = &dish.Detail

	return (*client).AddDish(*ctx, &req)
}

func CheckUserPassword(name, password string) (*msdbcall.CheckUserPasswordResp, error) {
	callback, client, ctx := GetMMDemoClient(define.MSDBCallIp + ":" + define.MSDBCallCheckUserPasswordPort)
	defer callback()

	req := msdbcall.CheckUserPasswordReq{}
	req.Name = &name
	req.Password = &password

	return (*client).CheckUserPassword(*ctx, &req)
}

func CheckUserNameUnique(name string) (*msdbcall.CheckUserNameUniqueResp, error) {
	callback, client, ctx := GetMMDemoClient(define.MSDBCallIp + ":" + define.MSDBCallCheckUserNameUniquePort)
	defer callback()

	req := msdbcall.CheckUserNameUniqueReq{}
	req.Name = &name

	return (*client).CheckUserNameUnique(*ctx, &req)
}
