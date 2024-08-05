package ero

import (
	"fmt"
	"runtime"
	"strings"
)

const unknownTrace = "unknown"

func getLineTrace() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return unknownTrace
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getStackTrace() (stacks []string) {
	for i := 3; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fun := runtime.FuncForPC(pc)
		funcPath := fun.Name()
		file = file[strings.LastIndexByte(file, '/')+1:]
		stacks = append(stacks, fmt.Sprintf("%s <%s:%d>", funcPath, file, line))
	}
	// hardcode remove last 2 basic levels stacktrace
	stacks = stacks[:len(stacks)-2]
	return
}
