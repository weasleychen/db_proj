package service

import (
	"db_proj/define"
	"db_proj/model"
	mstablemgr "db_proj/mstablemgr/proto"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

var (
	Tables = make([]mstablemgr.Table, 10)
	Mutex  sync.Mutex
	Times  atomic.Int32

	WALLogPath = fmt.Sprintf("%s/server/mstablemgr/wal.log", define.ProjectPath)
	WALLog, _  = os.OpenFile(WALLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	WALRedis   = model.WALRedisHandler.GetRedisClient()
)

func init() {
	Mutex.Lock()
	defer Mutex.Unlock()

	for i := 0; i < len(Tables); i++ {
		Tables[i] = mstablemgr.Table{
			Id:     int32(i + 1),
			Status: define.TableIsClosed,
		}
	}
}
