package main

import (
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
func RecoverFromLog(tables []mstablemgr.Table) {
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

			(&service.OpenTableServer{}).OpenTable(context.Background(), &req)
		} else if reqLog[0] == "CompleteTable" {
			req := mstablemgr.CompleteTableReq{}
			json.Unmarshal([]byte(reqLog[1]), &req)
			wal := false
			req.Wal = &wal

			(&service.CompleteTableServer{}).CompleteTable(context.Background(), &req)
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
	}
}

func init() {
	RecoverFromLog(service.Tables)
	go PersistTables()
}

func main() {
	go RegisterNewHandler(":"+define.MSTableMgrOpenTable, &service.OpenTableServer{})
	go RegisterNewHandler(":"+define.MSTableMgrCompleteTable, &service.CompleteTableServer{})
	go RegisterNewHandler(":"+define.MSTableMgrGetTablesStatus, &service.GetTablesStatusServer{})

	select {}
}
