package model

import (
	"math/big"
	"time"
)

type Block struct {
	Id           uint64
	ParentHash   string
	Hash         string
	Number       *big.Int
	Transactions []*EthTransaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
