package da

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockTx struct {
	mock.Mock
}

func (tx *MockTx) Commit() error {
	args := tx.Called()
	return args.Error(0)
}

func (tx *MockTx) Rollback() error {
	args := tx.Called()
	return args.Error(0)
}

func (tx *MockTx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	testArgs := tx.Called(ctx, query, args)
	return testArgs.Get(0).(*sql.Row)
}
func (tx *MockTx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	testArgs := tx.Called(ctx, query, args)
	return testArgs.Get(0).(sql.Result), testArgs.Error(1)
}
