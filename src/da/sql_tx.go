package da

import (
	"context"
	"database/sql"
)

type SqlTx interface {
	Commit() error
	Rollback() error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
