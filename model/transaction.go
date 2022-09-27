package model

import (
	"math/big"
)

type Transaction struct {
	ID          string
	BlockNumber *big.Int
	From        string
	To          string
	Amount      *big.Int
	Fee         *big.Int
	//Size   int
}

type EthTransaction struct {
	Transaction
	Gas                uint64
	GasPrice           *big.Int
	IsContractCreation bool
}
