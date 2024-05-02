package da

import (
	"context"
	"database/sql"
)

type SimpleWrapperSqlDb struct {
	db SqlDb
}

func (wdb *SimpleWrapperSqlDb) BeginTx(ctx context.Context, opts *sql.TxOptions) (SqlTx, error) {
	return wdb.db.BeginTx(ctx, opts)
}
func (wdb *SimpleWrapperSqlDb) Close() error {
	return wdb.db.Close()
}
