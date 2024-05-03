package error

import (
	"fmt"
	"strings"
)

// Errors raised by package x.
var (
	ErrDuplicate = wrapError{msg: "duplicate name"}
	ErrNotFound  = wrapError{msg: "not found"}
)

type wrapError struct {
	err error
	msg string
}

func (err wrapError) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s: %v", err.msg, err.err)
	}
	return err.msg
}
func (err wrapError) wrap(inner error) error {
	return wrapError{msg: err.msg, err: inner}
}
func (err wrapError) Unwrap() error {
	return err.err
}
func (err wrapError) Is(target error) bool {
	ts := target.Error()
	return ts == err.msg || strings.HasPrefix(ts, err.msg+": ")
}
