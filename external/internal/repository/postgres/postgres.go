package postgres

import "context"

type TransactionManager interface {
	GetQueryEngine(ctx context.Context) QueryEngine
	RunReadUncommitted(ctx context.Context, fn func(ctxTx context.Context) error) error
	RunSerializable(ctx context.Context, fn func(ctxTx context.Context) error) error
}

type PgRepository struct {
	txManager TransactionManager
}

func NewPgRepository(txManager TransactionManager) *PgRepository {
	return &PgRepository{
		txManager: txManager,
	}
}
