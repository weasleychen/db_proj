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

func GetMSTableMgrClient() (func(), *mstablemgr.MSTableMgrClient, *context.Context) {
	conn, err := grpc.Dial(
		define.MSTableMgrIP+":"+define.MSTableMgrPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial err")
	}

	mstablemgrclient := mstablemgr.NewMSTableMgrClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	callback := func() {
		conn.Close()
		cancel()
	}

	return callback, &mstablemgrclient, &ctx
}

func CallOpenTable(tableId int32) (*mstablemgr.OpenTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()
	wal := true
	return (*mstablemgrclient).OpenTable(*ctx, &mstablemgr.OpenTableReq{TableId: &tableId, Wal: &wal})
}

func CallCompleteTable(uin string, tableId int32) (*mstablemgr.CompleteTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	wal := true
	return (*mstablemgrclient).CompleteTable(*ctx, &mstablemgr.CompleteTableReq{TableId: &tableId, Uin: &uin, Wal: &wal})
}

func CallGetTablesStatus() (*mstablemgr.GetTablesStatusResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	return (*mstablemgrclient).GetTablesStatus(*ctx, &mstablemgr.GetTablesStatusReq{})
}

func CallAddTable(tableId int32) (*mstablemgr.AddTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	wal := true
	return (*mstablemgrclient).AddTable(*ctx, &mstablemgr.AddTableReq{TableId: &tableId, Wal: &wal})
}

func CallDelTable(tableId int32) (*mstablemgr.DelTableResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	wal := true
	return (*mstablemgrclient).DelTable(*ctx, &mstablemgr.DelTableReq{TableId: &tableId, Wal: &wal})
}

func CallOrderDish(tableId int32, dishIdList []int32) (*mstablemgr.OrderDishResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	wal := true
	return (*mstablemgrclient).OrderDish(*ctx, &mstablemgr.OrderDishReq{TableId: &tableId, DishIdList: dishIdList, Wal: &wal})
}

func CallGetTableInfo(tableId int32) (*mstablemgr.GetTableInfoResp, error) {
	callback, mstablemgrclient, ctx := GetMSTableMgrClient()
	defer callback()

	return (*mstablemgrclient).GetTableInfo(*ctx, &mstablemgr.GetTableInfoReq{TableId: &tableId})
}
