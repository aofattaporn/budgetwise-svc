package utils

import (
	"runtime"
	"strings"
)

func GetCurrentFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	arr := strings.Split(runtime.FuncForPC(pc[0]).Name(), ".")
	return arr[len(arr)-1]
}
