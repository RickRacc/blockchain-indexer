package repository

import (
	"context"
	"database/sql"
	"go-bonotans/model"
)

type TransactionPaymentRepository struct {
	dbPool *sql.DB
}

func NewTransactionPaymentRepository(pool *sql.DB) *TransactionPaymentRepository {
	return &TransactionPaymentRepository{
		dbPool: pool,
	}
}

func (repo *TransactionPaymentRepository) Save(ctx context.Context, block *model.TransactionPayment) error {
	return nil
}
