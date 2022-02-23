package sqli

import (
	"context"

	"github.com/jackc/pgx"
)

type PgxExecutor struct {
	*pgx.ConnPool
}

func NewQuerier(pool *pgx.ConnPool) *PgxExecutor {
	return &PgxExecutor{ConnPool: pool}
}

func (p *PgxExecutor) Executor() Executor {
	return p.ConnPool
}

func (p *PgxExecutor) WithTx(f func(executor Executor) error) error {
	return p.WithTx(f)
}

func (p *PgxExecutor) BeginEx(ctx context.Context, txOptions *pgx.TxOptions) (TxExecutor, error) {
	return p.BeginEx(ctx, txOptions)
}
