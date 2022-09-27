package model

import (
	"math/big"
	"time"
)

type Block struct {
	ParentHash   string
	Hash         string
	Number       *big.Int
	uint64       time.Time
	Transactions []*EthTransaction
}
