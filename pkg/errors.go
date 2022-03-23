package pkg

import (
	"fmt"
)

type MyError struct {
	Code 		int 	`json:"code"`
	Message 	string	`json:"message"`
}

func (e *MyError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Code)
}

func NotFoundError() *MyError {
	return &MyError{Code: 404, Message: "Not found"}
}

func UnauthorizedError() *MyError {
	return &MyError{Code: 401, Message: "Unauthorized"}
}

func NoContentError() *MyError {
	return &MyError{Code: 204, Message: "No content"}
}

func BadRequestError(msn string) *MyError {
	return &MyError{Code: 400, Message: msn}
}