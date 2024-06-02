package service

import (
	"db_proj/define"
	"db_proj/model"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

type Table struct {
	Id                int32        `json:"id,omitempty"`
	Status            int32        `json:"status,omitempty"`
	OrderedDishIdList []model.Dish `json:"orderedDishIdList"`
}

var (
	Tables = map[int]Table{}
	Mutex  sync.Mutex
	Times  atomic.Int32

	WALLogPath = fmt.Sprintf("%s/server/mstablemgr/wal.log", define.ProjectPath)
	WALLog, _  = os.OpenFile(WALLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	WALRedis   = model.WALRedisHandler.GetRedisClient()
)

func init() {
	Mutex.Lock()
	defer Mutex.Unlock()

	for i := 1; i <= 10; i++ {
		Tables[i] = Table{
			Id:                int32(i),
			Status:            define.TableIsNotInUse,
			OrderedDishIdList: make([]model.Dish, 0),
		}
	}
}
