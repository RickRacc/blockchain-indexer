package model

import (
	"time"
)

type Block struct {
	Id           uint64
	ParentHash   string
	Hash         string
	Number       int64
	Transactions []*EthTransaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
