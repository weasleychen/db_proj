package util

import (
	"db_proj/define"
	"fmt"
	"os"
	"time"
)

func Log(format string, args ...interface{}) {
	newFormat := fmt.Sprintf("[%s %s] %s\n", time.Now().Format(time.DateTime), GetStackInfo(2), format)
	file, _ := os.OpenFile(fmt.Sprintf("%s/server/log/%s.log", define.ProjectPath, time.Now().Format(time.DateOnly)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if len(args) == 0 {
		file.Write([]byte(fmt.Sprintf(newFormat)))
	} else {
		file.Write([]byte(fmt.Sprintf(newFormat, args...)))
	}
}
