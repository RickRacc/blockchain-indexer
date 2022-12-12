package test

import (
	"go-bonotans/model"
	"math/big"
)

func GetBlock() *model.Block {
	block := model.Block{
		ParentHash:   "parenthash",
		Hash:         "hash",
		Number:       1,
		Transactions: nil,
	}

	return &block
}

func GetEthTransaction() *model.EthTransaction {
	transaction := model.EthTransaction{
		BaseTransaction: model.BaseTransaction{
			Hash: "transactionHash",
			Fee:  big.NewInt(1000),
		},
		Gas:                big.NewInt(10),
		GasPrice:           big.NewInt(100),
		IsContractCreation: false,
	}

	return &transaction
}
