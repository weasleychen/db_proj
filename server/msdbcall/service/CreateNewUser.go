package service

import (
	"context"
	msdbcall "db_proj/msdbcall/proto"
)

type CreateNewUserServer struct {
	msdbcall.UnimplementedMSDBCallServer
}

func (server *CreateNewUserServer) CreateNewUser(ctx context.Context, req *msdbcall.CreateNewUserReq) (*msdbcall.CreateNewUserResp, error) {
	// TODO
	return nil, nil
}
