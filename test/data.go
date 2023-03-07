package test

import (
	"go-bonotans/model"
	"math/big"
)

var blocks = []model.Block{
	model.Block{
		ParentHash:   "parenthash",
		Hash:         "hash",
		Number:       1,
		Transactions: nil,
	},
	model.Block{
		ParentHash:   "parenthash2",
		Hash:         "hash2",
		Number:       2,
		Transactions: nil,
	},
}

func GetBlock() *model.Block {
	//block := model.Block{
	//	ParentHash:   "parenthash",
	//	Hash:         "hash",
	//	Number:       1,
	//	Transactions: nil,
	//}

	return &blocks[0]
}

func GetBlocks() []model.Block {
	//block := model.Block{
	//	ParentHash:   "parenthash",
	//	Hash:         "hash",
	//	Number:       1,
	//	Transactions: nil,
	//}

	return blocks
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

func GetTransactionPayment() *model.TransactionPayment {
	transactionPayment := model.TransactionPayment{
		From:   "from",
		To:     "to",
		Index:  0,
		Amount: big.NewInt(1000),
	}

	return &transactionPayment
}
