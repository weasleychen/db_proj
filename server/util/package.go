package util

import (
	"fmt"
	"runtime"
)

// 调用堆栈形式，跳跃skip层，获得模块完整import路径与函数名与行号(调用者, <import_path>.function:line_number)
func GetStackInfo(skip int) string {
	pc, _, line, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name() + fmt.Sprintf(":%d", line)
}
