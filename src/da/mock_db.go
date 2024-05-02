package da

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDb struct {
	mock.Mock
}

func (db *MockDb) BeginTx(ctx context.Context, opts *sql.TxOptions) (SqlTx, error) {
	args := db.Called(ctx, opts)
	return args.Get(0).(SqlTx), args.Error(1)
}
func (db *MockDb) Close() error {
	args := db.Called()
	return args.Error(0)
}
