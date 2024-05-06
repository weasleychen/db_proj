package mstablemgrclient

import (
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func GetMSTableMgrClient(address string) (func(), *mstablemgr.MSTableMgrClient, *context.Context) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err")
	}

	mstablemgrclient := mstablemgr.NewMSTableMgrClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	callback := func() {
		conn.Close()
		cancel()
	}

	return callback, &mstablemgrclient, &ctx
}

func CallOpenTable(tableId int32) (*mstablemgr.OpenTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient(define.MSTableMgrIP + ":" + define.MSTableMgrOpenTable)
	defer callback()
	wal := true
	return (*mstablemgrclient).OpenTable(*ctx, &mstablemgr.OpenTableReq{TableId: &tableId, Wal: &wal})
}

func CallCompleteTable(tableId int32) (*mstablemgr.CompleteTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient(define.MSTableMgrIP + ":" + define.MSTableMgrCompleteTable)
	defer callback()

	wal := true
	return (*mstablemgrclient).CompleteTable(*ctx, &mstablemgr.CompleteTableReq{TableId: &tableId, Wal: &wal})
}

func CallGetTablesStatus() (*mstablemgr.GetTablesStatusResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient(define.MSTableMgrIP + ":" + define.MSTableMgrGetTablesStatus)
	defer callback()

	return (*mstablemgrclient).GetTablesStatus(*ctx, &mstablemgr.GetTablesStatusReq{})
}
