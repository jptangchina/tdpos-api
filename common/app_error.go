package common

import "net/http"

type AppError struct {
	HttpStatus int
	Response   *Response
}

func NewAppError(appStatus int, msg string) *AppError {
	return &AppError{http.StatusOK, &Response{appStatus, msg, nil}}
}
