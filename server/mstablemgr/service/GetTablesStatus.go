package service

import (
	"context"
	mstablemgr "db_proj/mstablemgr/proto"
	"sort"
)

type GetTablesStatusServer struct {
	mstablemgr.UnimplementedMSTableMgrServer
}

func (server *GetTablesStatusServer) GetTablesStatus(ctx context.Context, req *mstablemgr.GetTablesStatusReq) (*mstablemgr.GetTablesStatusResp, error) {
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
