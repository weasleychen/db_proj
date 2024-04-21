package util

import (
	"fmt"
	"runtime"
)

func GetFunctionName(depth int) string {
	pc, _, line, _ := runtime.Caller(depth)
	return runtime.FuncForPC(pc).Name() + fmt.Sprintf(":%d", line)
}
