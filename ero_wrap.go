package ero

import (
	"errors"
	"fmt"
)

func (eme *eroMsgError) Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &eroMsgError{
		msg: text,
		lt:  getLineTrace(),
		c:   err,
	}
}

func Wrap(err error, text string, args ...any) error {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}
wrap:
	e, ok := as[interface{ Wrap(error, string) error }](err)
	if !ok {
		if err == nil {
			err = errors.New("nil error")
		} else {
			err = newEroError(err.Error())
		}
		goto wrap
	}
	return e.Wrap(err, text)
}

func (eme *eroMsgError) UnwrapOnce(err error) error {
	if err == nil {
		return nil
	}
	e, ok := as[interface{ Child() error }](err)
	if !ok {
		return err
	}
	return e.Child()
}

func UnwrapOnce(err error) error {
	e := mustAs[interface{ UnwrapOnce(error) error }](err)
	return e.UnwrapOnce(err)
}

func UnwrapAll(err error) (errs []error) {
	if err == nil {
		return
	}
	walkUnwrap(err, func(err error, _ bool) bool {
		errs = append(errs, err)
		return true
	})
	return errs
}

func walkUnwrap(err error, fun func(err error, isLast bool) bool) {
	if err != nil {
		e := mustAs[interface{ Child() error }](err)
		child := e.Child()
		isLast := child == nil
		if !fun(err, isLast) {
			return
		}
		walkUnwrap(child, fun)
	}
}
