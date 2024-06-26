package service

import (
	"context"
	"db_proj/define"
	"db_proj/model"
	msdbcallclient "db_proj/msdbcall/client"
	mstablemgr "db_proj/mstablemgr/proto"
	"db_proj/util"
	"encoding/json"
	"fmt"
	"sort"
	"time"
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

	Tables[int(*req.TableId)] = Table{
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
		getUserDiscountResp, err := msdbcallclient.CallGetUserDiscount(req.GetUin())
		if err != nil {
			return nil, err
		}

		originPrice := 0.0
		for _, dish := range table.OrderedDishIdList {
			originPrice += dish.Price * dish.Discount
			util.Log("dish.price: %v, discount: %v", dish.Price, dish.Discount)
		}

		recordJson := util.MarshalJson(&model.ConsumeRecordJson{
			TableId:       table.Id,
			Uin:           req.GetUin(),
			OrderedDishes: table.OrderedDishIdList,
			Discount:      getUserDiscountResp.GetDiscount(),
			OriginPrice:   originPrice,
			FinalPrice:    originPrice * getUserDiscountResp.GetDiscount(),
			CompleteTableTime: time.Now().Unix(),		
		})

		_, err = msdbcallclient.CallStoreConsumeRecord(recordJson)
		if err != nil {
			return nil, err
		}

		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("CompleteTable %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	Tables[int(req.GetTableId())] = Table{
		Id:                req.GetTableId(),
		Status:            define.TableIsNotInUse,
		OrderedDishIdList: make([]model.Dish, 0),
	}
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

func (server *MSTableMgrServer) OrderDish(ctx context.Context, req *mstablemgr.OrderDishReq) (*mstablemgr.OrderDishResp, error) {
	fmt.Printf("Order Dish\n")
	Mutex.Lock()
	defer Mutex.Unlock()

	resp := mstablemgr.OrderDishResp{}

	tableId := int(req.GetTableId())

	if Tables[tableId].Status == define.TableIsNotInUse {
		resp.Status = util.NewType[int32](define.TableIsNotInUse)
		return &resp, nil
	}

	isSomeDishNotExist := false
	dishes := make([]model.Dish, 0)
	for _, dishId := range req.GetDishIdList() {
		getDishInfoResp, err := msdbcallclient.CallGetDishInfo(dishId)
		if err != nil {
			return nil, err
		}

		if getDishInfoResp.GetStatus() == define.ErrorDishIdNotExist {
			isSomeDishNotExist = true
			resp.NotExistDish = append(resp.NotExistDish, dishId)
		} else {
			modelDish := model.Dish{}
			modelDish.ID = uint(getDishInfoResp.GetDish().GetId())
			modelDish.Name = getDishInfoResp.GetDish().GetName()
			modelDish.Price = getDishInfoResp.GetDish().GetPrice()
			modelDish.Discount = getDishInfoResp.GetDish().GetDiscount()
			modelDish.Detail = getDishInfoResp.GetDish().GetDetail()

			dishes = append(dishes, modelDish)
		}
	}

	if isSomeDishNotExist {
		resp.Status = util.NewType[int32](define.ErrorDishIdNotExist)
		return &resp, nil
	}

	table := Tables[tableId]
	table.OrderedDishIdList = append(table.OrderedDishIdList, dishes...)
	Tables[tableId] = table

	if req.GetWal() {
		jsonBytes, _ := json.Marshal(req)
		WALLog.Write([]byte(fmt.Sprintf("OrderDish %s\n", string(jsonBytes))))
		Times.Add(1)
	}

	resp.Status = util.NewType[int32](define.OK)
	return &resp, nil
}

func (server *MSTableMgrServer) GetTableInfo(ctx context.Context, req *mstablemgr.GetTableInfoReq) (*mstablemgr.GetTableInfoResp, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	resp := mstablemgr.GetTableInfoResp{}

	tableId := int(req.GetTableId())

	table := Tables[tableId]
	jsonString := util.MarshalJson(table)

	resp.JsonString = &jsonString
	resp.Status = util.NewType[int32](define.OK)

	return &resp, nil
}
