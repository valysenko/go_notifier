package errors

import (
	"fmt"
)

const (
	NoType = ErrorType(iota)
	RequestErrCopyToBuffer
	RequestErrUnmarshal
	RequestErrNotValid
)

type ErrorType uint

type WrappedError struct {
	ErrorType ErrorType
	Context   string
	Err       error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string, errorType ErrorType) *WrappedError {
	return &WrappedError{
		Context:   info,
		Err:       err,
		ErrorType: errorType,
	}
}
