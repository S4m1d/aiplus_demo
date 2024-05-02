package da

import (
	"context"
	"database/sql"
)

type WrapperSqlDb interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SqlTx, error)
	Close() error
}
