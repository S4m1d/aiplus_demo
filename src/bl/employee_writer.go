package bl

import (
	"aiplus_demo/src/da"
	"context"
)

type EmployeeWriter interface {
	Write(ctx context.Context, empl da.Employee) (bool, *EmployeeWriteErr)
}

type EmployeeWriteErr struct {
	ErrType BlErrorType
	message string
}

func NewEmployeeWriteErr(errType BlErrorType, message string) *EmployeeWriteErr {
	return &EmployeeWriteErr{
		ErrType: errType,
		message: message,
	}
}

func (e EmployeeWriteErr) Error() string {
	return e.message
}
