package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
	"math/big"
)

type TransactionPaymentRepository struct {
	pool *sql.DB
}

func NewTransactionPaymentRepository(pool *sql.DB) *TransactionPaymentRepository {
	return &TransactionPaymentRepository{
		pool: pool,
	}
}

func (repo *TransactionPaymentRepository) Save(ctx context.Context, payment *model.TransactionPayment) (*model.TransactionPayment, error) {
	stmt := fmt.Sprintf("insert into transaction_payment (%s) values ('%v', '%s', '%s', '%d', '%s') returning %s",
		TRANSACTION_PAYMENT_INSERT_COLS, payment.TransactionId, payment.From, payment.To, payment.Index,
		payment.Amount, TRANSACTION_PAYMENT_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *TransactionPaymentRepository) Read(row *sql.Row) (*model.TransactionPayment, error) {
	var amount []byte
	var payment model.TransactionPayment

	err := row.Scan(&payment.Id, &payment.TransactionId, &payment.From, &payment.To,
		&payment.Index, &amount, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	payment.Amount = new(big.Int).SetBytes(amount)

	return &payment, nil
}
