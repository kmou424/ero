package ero

import (
	"fmt"
	"strings"
)

func LineTrace(err error) string {
	if err == nil {
		return fmt.Sprint(err)
	}
	e := mustAs[interface{ This() *eroMsgError }](err).This()
	return fmt.Sprintf("Error: %s\n  => at %s", e.msg, e.lt)
}

func unwrapError(sb *strings.Builder, err error) error {
	isChild := false
	var lastErr error

	childTrace := func(err error) string {
		if err == nil {
			return "<nil>"
		}
		e := mustAs[interface{ This() *eroMsgError }](err).This()
		return fmt.Sprintf("Caused by: %s\n  => at %s", e.msg, e.lt)
	}

	walkUnwrap(err, func(e error, last bool) bool {
		if isChild {
			sb.WriteString(childTrace(e))
		} else {
			sb.WriteString(LineTrace(e))
			isChild = true
		}
		if last {
			lastErr = e
			sb.WriteString(" [source]")
		} else {
			sb.WriteString(" [wrapped]")
			sb.WriteString("\n")
		}
		return true
	})
	return lastErr
}

func AllTrace(err error, showStackTrace ...bool) string {
	if err == nil {
		return fmt.Sprint(err)
	}
	var sb strings.Builder
	err = unwrapError(&sb, err)
	if len(showStackTrace) > 0 && showStackTrace[0] {
		sb.WriteRune('\n')
		sb.WriteRune('\n')
		sb.WriteString(StackTrace(err))
	}
	return sb.String()
}

func StackTrace(err error) string {
	if err == nil {
		return fmt.Sprint(err)
	}
	for { // iterate over all wrapped errors
		child := UnwrapOnce(err)
		if child == nil {
			break
		}
		err = child
	}
	var sb strings.Builder
	sb.WriteString("Source Stacktrace:")
	if e, ok := as[interface{ StackTrace() []string }](err); ok {
		for _, s := range e.StackTrace() {
			sb.WriteRune('\n')
			sb.WriteString("  ")
			sb.WriteString(s)
		}
	}
	return sb.String()
}
