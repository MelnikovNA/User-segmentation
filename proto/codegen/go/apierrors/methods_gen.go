package apierrors

import (
	"fmt"
	"regexp"
)

const (
	InternalServerErrorName = "InternalServerError"
	InternalServerErrorCode = 500
)

type APIError struct {
	name    string
	code    int
	message string
}

func (a *APIError) String() string {
	return a.message
}

func (a *APIError) Error() error {
	return fmt.Errorf("{%v}%v", a.name, a.message)
}

var re = regexp.MustCompile(`{(\w*)}`)

func FromError(err error) *APIError {
	matches := re.FindStringSubmatch(err.Error())
	key := ""
	if len(matches) >= 2 {
		key = matches[1]
	}
	res, ok := apiErrors[key]
	if !ok {
		e := &APIError{
			name:    InternalServerErrorName,
			code:    InternalServerErrorCode,
			message: fmt.Sprintf("no matched error:%v", err),
		}
		return e
	}
	return res
}

func (a *APIError) Name() string {
	return a.name
}

func (a *APIError) Message() string {
	return a.message
}

func (a *APIError) Code() int {
	return a.code
}
