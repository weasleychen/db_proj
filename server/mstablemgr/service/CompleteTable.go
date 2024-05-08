package service

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"encoding/json"
	"fmt"
)

type CompleteTableServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *CompleteTableServer) LogToWAL(req *mstablemgr.CompleteTableReq) {
	jsonBytes, _ := json.Marshal(req)
	WALLog.Write([]byte(fmt.Sprintf("CompleteTable %s\n", string(jsonBytes))))
}

func (server *CompleteTableServer) CompleteTable(ctx context.Context, req *mstablemgr.CompleteTableReq) (*mstablemgr.CompleteTableResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	var status int32
	resp := mstablemgr.CompleteTableResp{Status: &status}

	table, exist := Tables[int(*req.TableId)]
	if !exist {
		status = define.ErrorTableIdNotExist
		return &resp, nil
	}

	if table.Status == define.TableIsNotInUse {
		status = define.ErrorTableIsClosed
		return &resp, nil
	}

	if *req.Wal {
		server.LogToWAL(req)
		Times.Add(1)
	}

	table.Status = define.TableIsNotInUse
	Tables[int(*req.TableId)] = table
	status = define.OK

	return &resp, nil
}
