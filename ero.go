package ero

import (
	"fmt"
)

type eroMsgError struct {
	msg string
	lt  string
	ss  []string
	c   error
}

func (eme *eroMsgError) This() *eroMsgError {
	return eme
}

func (eme *eroMsgError) Child() error {
	return eme.c
}

func (eme *eroMsgError) StackTrace() (stackTrace []string) {
	if eme == nil {
		return
	}
	return eme.ss
}

func (eme *eroMsgError) Is(err error) bool {
	if err == nil {
		return false
	}
	if EqualPtr(eme, err) {
		return true
	}
	return eme.Error() == err.Error()
}

func (eme *eroMsgError) Error() string {
	return eme.msg
}

func newEroError(s string) error {
	return &eroMsgError{
		msg: s,
		lt:  getLineTrace(),
		ss:  getStackTrace(),
	}
}

func New(text string) error {
	return newEroError(text)
}

func Newf(format string, args ...any) error {
	return newEroError(fmt.Sprintf(format, args...))
}

func Default() error {
	return newEroError("an error occurred here")
}
