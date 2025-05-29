package ero

import "sync/atomic"

type TraceLevel int32

const (
	TraceLevelNone TraceLevel = iota
	TraceLevelLine
	TraceLevelFull
	TraceLevelFullWithStack
)

var traceLevel = TraceLevelNone

func SetTraceLevel(level TraceLevel) {
	atomic.StoreInt32((*int32)(&traceLevel), int32(level))
}

func GetTraceLevel() TraceLevel {
	return TraceLevel(atomic.LoadInt32((*int32)(&traceLevel)))
}
