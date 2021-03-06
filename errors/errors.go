package errors

import (
	"fmt"
	"strings"
	"time"
)

// Error is medtune implementation of storing
// multiple error informations in one object
type Error struct {
	Code int64 `json:"code"`
	// Main message
	Message string `json:"message"`
	// Type
	Type string `json:"type"`
	// Occure type
	Time time.Time `json:"time"`
	// Sub errors
	SubErrors []error `json:"sub_errors"`
}

func errorsToStrings(errs ...error) []string {
	s := make([]string, 0, len(errs))
	for _, err := range errs {
		s = append(s, err.Error())
	}
	return s
}

// Error implementing error interface
func (e *Error) Error() string {
	if len(e.SubErrors) == 0 {
		return fmt.Sprintf("%s", e.Message)
	}
	return fmt.Sprintf("%s\n\t%v\n", e.Message, strings.Join(errorsToStrings(e.SubErrors...), "\n\t"))
}

// Append adds errors
func (e *Error) Append(err ...error) {
	e.SubErrors = append(e.SubErrors, err...)
}

// Appendf adds errors
func (e *Error) Appendf(format string, a ...interface{}) {
	e.SubErrors = append(e.SubErrors, fmt.Errorf(format, a...))
}

// IsEmpty checks if error is empty
func (e *Error) IsEmpty() bool {
	return e.Message == "" && len(e.SubErrors) == 0
}

// Nil return an empty (represent nil) error
func Nil() *Error {
	return &Error{}
}

// New returns a new simple error
func New(err string) *Error {
	return &Error{
		Message: err,
		Time:    time.Now(),
	}
}

// NewList returns a list of errors as sub errors
func NewList(errors ...error) *Error {
	return &Error{
		Time:      time.Now(),
		SubErrors: errors,
	}
}

// Newf equivalent to fmt.Errorf
func Newf(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
		Time:    time.Now(),
	}
}
