package main

import (
	"fmt"
	"bufio"
	"context"
	"db_proj/define"
	mstablemgr "db_proj/mstablemgr/proto"
	"db_proj/mstablemgr/service"
	"db_proj/util"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

func RegisterNewHandler(address string, handler mstablemgr.MSTableMgrServer) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("listen on ", address, " error")
		return err
	}

	server := grpc.NewServer()
	mstablemgr.RegisterMSTableMgrServer(server, handler)

	return server.Serve(listener)
}

// RecoverFromLog WAL-从日志中恢复数据
func RecoverFromLog(tables map[int]service.Table) {
	walJsonBytes, err := service.WALRedis.Get(define.DefaultRedisContext, "log").Bytes()
	if err != nil && !errors.Is(err, redis.Nil) {
		util.Log("WAL read Redis error: ", err)
		return
	}

	if !errors.Is(err, redis.Nil) {
		json.Unmarshal(walJsonBytes, &tables)
	}

	scanner := bufio.NewScanner(service.WALLog)
	for scanner.Scan() {
		line := scanner.Text()

		util.Log("recover: %s", line)
		reqLog := strings.Split(line, " ")

		if reqLog[0] == "OpenTable" {
			req := mstablemgr.OpenTableReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			resp, err := (&service.MSTableMgrServer{}).OpenTable(context.Background(), &req)
			util.Log("log: %v, resp: %v, err: %v", reqLog, resp, err)
		} else if reqLog[0] == "CompleteTable" {
			req := mstablemgr.CompleteTableReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			resp, err := (&service.MSTableMgrServer{}).CompleteTable(context.Background(), &req)
			util.Log("log: %v, resp: %v, err: %v", reqLog, resp, err)
		} else if reqLog[0] == "AddTable" {
			req := mstablemgr.AddTableReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			resp, err := (&service.MSTableMgrServer{}).AddTable(context.Background(), &req)
			util.Log("log: %v, resp: %v, err: %v", reqLog, resp, err)
		} else if reqLog[0] == "DelTable" {
			req := mstablemgr.DelTableReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			resp, err := (&service.MSTableMgrServer{}).DelTable(context.Background(), &req)
			util.Log("log: %v, resp: %v, err: %v", reqLog, resp, err)
		} else if reqLog[0] == "OrderDish" {
			req := mstablemgr.OrderDishReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			resp, err := (&service.MSTableMgrServer{}).OrderDish(context.Background(), &req)
			util.Log("log: %v, resp: %v, err: %v", reqLog, resp, err)
		} else {
			fmt.Printf("unexpected branch: %v\n", reqLog)
		}
	}
}

func PersistTables() {
	for {
		if service.Times.Load() == 50 {
			jsonBytes, err := json.Marshal(service.Tables)
			if err != nil {
				util.Log("Marshal tables error: %v", err)
				continue
			}

			service.WALRedis.Set(define.DefaultRedisContext, "log", jsonBytes, 0)

			service.WALLog.Truncate(0)
			service.Times.Store(0)
		}
		time.Sleep(1 * time.Second)
	}
}

func init() {
	time.Sleep(2 * time.Second)
	RecoverFromLog(service.Tables)
	go PersistTables()
}

func main() {
	RegisterNewHandler(define.MSTableMgrIP+":"+define.MSTableMgrPort, &service.MSTableMgrServer{})
}
