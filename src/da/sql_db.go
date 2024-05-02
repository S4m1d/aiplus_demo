package da

import (
	"context"
	"database/sql"
)

type SqlDb interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}
