package repository

import (
	"context"
	"database/sql"
	"go-bonotans/model"
)

type TransactionRepository struct {
	pool *sql.DB
}

func NewTransactionRepository(pool *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		pool: pool,
	}
}

func (repo *TransactionRepository) Save(ctx context.Context, block *model.EthTransaction) error {
	return nil
}
