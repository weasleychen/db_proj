package service

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"encoding/json"
	"fmt"
)

type OpenTableServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *OpenTableServer) LogToWAL(req *mstablemgr.OpenTableReq) {
	jsonBytes, _ := json.Marshal(req)
	WALLog.Write([]byte(fmt.Sprintf("OpenTable %s\n", string(jsonBytes))))
}

func (server *OpenTableServer) OpenTable(ctx context.Context, req *mstablemgr.OpenTableReq) (*mstablemgr.OpenTableResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	var status int32
	resp := mstablemgr.OpenTableResp{Status: &status}

	table, exist := Tables[int(*req.TableId)]
	if !exist {
		status = define.ErrorTableIdNotExist
		return &resp, nil
	}

	if table.Status == define.TableIsInUse {
		status = define.ErrorTableIsOpened
		return &resp, nil
	}

	if *req.Wal {
		server.LogToWAL(req)
		Times.Add(1)
	}

	table.Status = define.TableIsInUse
	Tables[int(*req.TableId)] = table
	status = define.OK

	return &resp, nil
}
