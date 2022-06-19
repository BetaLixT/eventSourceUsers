package blerr

import "fmt"

type Error struct {
	StatusCode   int    `json:"-"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorDetail  string `json:"errorDetail,omitempty"`
}

func (err *Error) Error() string {
	return fmt.Sprintf(
		"%d [%s]: %s | Requested Status Code: %d",
		err.ErrorCode,
		err.ErrorMessage,
		err.ErrorDetail,
		err.StatusCode,
	)
}

var _ error = (*Error)(nil)

func NewError(code ErrorCode, status int, detail string) *Error {
	return &Error{
		StatusCode: status,
		ErrorCode: code.code,
		ErrorMessage: code.message,
		ErrorDetail: detail,
	}
}

func UnexpectedError() *Error {
	return &Error{
		StatusCode:   500,
		ErrorCode:    UnexpectedErrorCode.code,
		ErrorMessage: UnexpectedErrorCode.message,
	}
}

func UnsetResponseError() *Error {
	return &Error{
		StatusCode: 500,
		ErrorCode: UnsetResponse.code,
		ErrorMessage: UnsetResponse.message,
	}
}

func UnexpectedErrorDetailed(e error) *Error {
	return &Error{
		StatusCode:   500,
		ErrorCode:    UnexpectedErrorCode.code,
		ErrorMessage: UnexpectedErrorCode.message,
		ErrorDetail:  e.Error(),
	}
}

func NotImplemented() *Error {
	return &Error{
		StatusCode:   501,
		ErrorCode:    NotImplementedCode.code,
		ErrorMessage: NotImplementedCode.message,
	}
}
