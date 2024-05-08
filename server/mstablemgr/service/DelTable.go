package service

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"encoding/json"
	"fmt"
)

type DelTableServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *DelTableServer) LogToWAL(req *mstablemgr.DelTableReq) {
	jsonBytes, _ := json.Marshal(req)
	WALLog.Write([]byte(fmt.Sprintf("DelTable %s\n", string(jsonBytes))))
}

func (server *DelTableServer) DelTable(ctx context.Context, req *mstablemgr.DelTableReq) (*mstablemgr.DelTableResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	var status int32
	resp := mstablemgr.DelTableResp{Status: &status}

	_, exist := Tables[int(*req.TableId)]
	if !exist {
		status = define.ErrorTableIdNotExist
		return &resp, nil
	}

	if *req.Wal {
		server.LogToWAL(req)
		Times.Add(1)
	}

	delete(Tables, int(*req.TableId))
	status = define.OK
	return &resp, nil
}
