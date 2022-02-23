package sqli

import (
	"context"
	"io"

	"github.com/jackc/pgx"
)

type Executor interface {
	ExecEx(ctx context.Context, sql string, options *pgx.QueryExOptions, arguments ...interface{}) (commandTag pgx.CommandTag, err error)
	QueryEx(ctx context.Context, sql string, options *pgx.QueryExOptions, args ...interface{}) (*pgx.Rows, error)
	QueryRowEx(ctx context.Context, sql string, options *pgx.QueryExOptions, args ...interface{}) *pgx.Row
	PrepareEx(ctx context.Context, name, sql string, opts *pgx.PrepareExOptions) (*pgx.PreparedStatement, error)
	CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int, error)
	CopyFromReader(r io.Reader, sql string) (pgx.CommandTag, error)
	CopyToWriter(w io.Writer, sql string, args ...interface{}) (pgx.CommandTag, error)
}

type TxExecutor interface {
	Executor
	CommitEx(ctx context.Context) error
	RollbackEx(ctx context.Context) error
}

type Querier interface {
	Executor() Executor
	WithTx(func(executor Executor) error) error
	BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (TxExecutor, error)
}
