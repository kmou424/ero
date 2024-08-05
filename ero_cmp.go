package ero

import (
	"unsafe"
)

func EqualPtr(err error, target error) bool {
	return (*[2]uintptr)(unsafe.Pointer(&err))[1] == (*[2]uintptr)(unsafe.Pointer(&target))[1]
}

func Equal(err error, target error) (isEqual bool) {
	if err == nil || target == nil {
		return EqualPtr(err, target)
	}
	walkUnwrap(err, func(innerErr error, _ bool) bool {
		switch x := err.(type) {
		case interface{ Is(error) bool }:
			if x.Is(target) {
				isEqual = true
				return false
			}
		default:
			if EqualPtr(err, target) {
				isEqual = true
				return false
			}
		}
		return true
	})
	return
}
