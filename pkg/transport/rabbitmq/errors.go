package rabbitmq

import "fmt"

const (
	NoType = ErrorType(iota)
	Retriable
	Skippable
)

type ErrorType uint

type HandlerError struct {
	ErrorType ErrorType
	Context   string
	Err       error
}

func (he *HandlerError) Error() string {
	return fmt.Sprintf("%s: %v", he.Context, he.Err)
}

func NewRetriableError(err error, info string) *HandlerError {
	return &HandlerError{
		Context:   info,
		Err:       err,
		ErrorType: Retriable,
	}
}

func NewSkippableError(err error, info string) *HandlerError {
	return &HandlerError{
		Context:   info,
		Err:       err,
		ErrorType: Skippable,
	}
}
