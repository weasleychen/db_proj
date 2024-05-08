package service

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"encoding/json"
	"fmt"
)

type AddTableServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *AddTableServer) LogToWAL(req *mstablemgr.AddTableReq) {
	jsonBytes, _ := json.Marshal(req)
	WALLog.Write([]byte(fmt.Sprintf("AddTable %s\n", string(jsonBytes))))
}

func (server *AddTableServer) AddTable(ctx context.Context, req *mstablemgr.AddTableReq) (*mstablemgr.AddTableResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	var status int32
	resp := mstablemgr.AddTableResp{Status: &status}

	_, exist := Tables[int(*req.TableId)]
	if exist {
		status = define.ErrorTableIdExist
		return &resp, nil
	}

	Tables[int(*req.TableId)] = mstablemgr.Table{
		Id:     *req.TableId,
		Status: define.TableIsNotInUse,
	}

	if *req.Wal {
		server.LogToWAL(req)
		Times.Add(1)
	}

	status = define.OK
	return &resp, nil
}
