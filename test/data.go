package test

import (
	"go-bonotans/model"
	"math/big"
)

func GetBlock() *model.Block {
	block := model.Block{
		ParentHash:   "parenthash",
		Hash:         "hash",
		Number:       new(big.Int).SetInt64(1),
		Transactions: nil,
	}

	return &block
}
