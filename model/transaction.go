package model

import (
	"math/big"
	"time"
)

type Transaction interface {
	//Chain() string
	Chain() any
}

type BaseTransaction struct {
	Id          uint64
	Hash        string
	BlockNumber *big.Int
	Fee         *big.Int
	Payments    []*TransactionPayment
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *BaseTransaction) Chain() string {
	return ""
}

type EthTransaction struct {
	BaseTransaction
	Gas                *big.Int
	GasPrice           *big.Int
	IsContractCreation bool
}

func (t *EthTransaction) Chain() string {
	return "ETH"
}
