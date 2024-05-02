package da

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockEmployeeRepository struct {
	mock.Mock
}

func (r *MockEmployeeRepository) IsEmployeeWithNumberExist(ctx context.Context, phoneNumber string, tx SqlTx) (bool, error) {
	args := r.Called(ctx, phoneNumber, tx)
	return args.Bool(0), args.Error(1)
}
func (r *MockEmployeeRepository) CreateEmployee(ctx context.Context, empl Employee, tx SqlTx) error {
	args := r.Called(ctx, empl, tx)
	return args.Error(0)
}
