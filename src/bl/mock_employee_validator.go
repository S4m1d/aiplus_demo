package bl

import (
	"aiplus_demo/src/da"

	"github.com/stretchr/testify/mock"
)

type MockEmployeeValidator struct {
	mock.Mock
}

func (ew *MockEmployeeValidator) Validate(empl da.Employee) error {
	args := ew.Called(empl)
	return args.Error(0)
}
