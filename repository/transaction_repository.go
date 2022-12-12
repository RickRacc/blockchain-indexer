package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
	"math/big"
)

type TransactionRepository struct {
	pool *sql.DB
}

func NewTransactionRepository(pool *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		pool: pool,
	}
}

func (repo *TransactionRepository) Save(ctx context.Context, transaction *model.EthTransaction) (*model.EthTransaction, error) {
	stmt := fmt.Sprintf("insert into eth_transaction (%s) values ('%s', '%v', %s, %s, %s, %t) returning %s",
		TRANSACTION_INSERT_COLS, transaction.Hash, transaction.BlockNumber, transaction.Fee,
		transaction.Gas, transaction.GasPrice, transaction.IsContractCreation, TRANSACTION_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	t, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (repo *TransactionRepository) Read(row *sql.Row) (*model.EthTransaction, error) {
	var fee []byte
	var gas []byte
	var gasPrice []byte

	var transaction model.EthTransaction

	err := row.Scan(&transaction.Id, &transaction.Hash, &transaction.BlockNumber, &fee, &gas, &gasPrice,
		&transaction.IsContractCreation, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return nil, err
	}
	transaction.Fee = new(big.Int).SetBytes(fee)
	transaction.Gas = new(big.Int).SetBytes(gas)
	transaction.GasPrice = new(big.Int).SetBytes(gasPrice)

	return &transaction, nil
}
