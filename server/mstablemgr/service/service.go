package service

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"encoding/json"
	"fmt"
	"sort"
)

type MSTableMgrServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *MSTableMgrServer) AddTable(ctx context.Context, req *mstablemgr.AddTableReq) (*mstablemgr.AddTableResp, error) {
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
		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("AddTable %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	status = define.OK
	return &resp, nil
}

func (server *MSTableMgrServer) CompleteTable(ctx context.Context, req *mstablemgr.CompleteTableReq) (*mstablemgr.CompleteTableResp, error) {
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
		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("CompleteTable %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	table.Status = define.TableIsNotInUse
	Tables[int(*req.TableId)] = table
	status = define.OK

	return &resp, nil
}

func (server *MSTableMgrServer) DelTable(ctx context.Context, req *mstablemgr.DelTableReq) (*mstablemgr.DelTableResp, error) {
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
		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("DelTable %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	delete(Tables, int(*req.TableId))
	status = define.OK
	return &resp, nil
}

func (server *MSTableMgrServer) GetTablesStatus(ctx context.Context, req *mstablemgr.GetTablesStatusReq) (*mstablemgr.GetTablesStatusResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	resp := mstablemgr.GetTablesStatusResp{}

	for _, table := range Tables {
		resp.TableList = append(resp.TableList, &mstablemgr.TableInfo{
			TableId: &table.Id,
			Status:  &table.Status,
		})
	}

	sort.Slice(resp.TableList, func(lhs, rhs int) bool {
		return *resp.TableList[lhs].TableId < *resp.TableList[rhs].TableId
	})

	return &resp, nil
}

func (server *MSTableMgrServer) OpenTable(ctx context.Context, req *mstablemgr.OpenTableReq) (*mstablemgr.OpenTableResp, error) {
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
		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("OpenTable %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	table.Status = define.TableIsInUse
	Tables[int(*req.TableId)] = table
	status = define.OK

	return &resp, nil
}
