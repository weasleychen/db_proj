package util

import "time"

/*
	高50位时间，低14位id
*/

var (
	LastTime int64 = 0
	id       int64 = 0
)

// 生成唯一uin，50位时间+14位id 并md5加密
func GenNewUin() string {
	if now := time.Now().Unix(); now != LastTime {
		LastTime = now
		id = 0
	}

	id += 1
	return EncryptInt64((LastTime << 50) | id)
}
