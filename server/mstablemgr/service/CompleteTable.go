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

	if *req.TableId < 1 || int(*req.TableId) > len(Tables) {
		status = define.ErrorTableIdInvalid
		return &resp, nil
	}

	if Tables[*req.TableId-1].Status == define.TableIsClosed {
		status = define.ErrorTableIsClosed
		return &resp, nil
	}

	if *req.Wal {
		server.LogToWAL(req)
	}
	Times.Add(1)

	Tables[*req.TableId-1].Status = define.TableIsClosed
	status = define.OK

	return &resp, nil
}
