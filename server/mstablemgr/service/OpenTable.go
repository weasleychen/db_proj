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

	if *req.TableId < 1 || int(*req.TableId) > len(Tables) {
		status = define.ErrorTableIdInvalid
		return &resp, nil
	}

	if Tables[*req.TableId-1].Status == define.TableIsOpened {
		status = define.ErrorTableIsOpened
		return &resp, nil
	}

	if *req.Wal {
		server.LogToWAL(req)
	}
	Times.Add(1)

	Tables[*req.TableId-1].Status = define.TableIsOpened
	status = define.OK

	return &resp, nil
}
