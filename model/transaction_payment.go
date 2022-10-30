package model

import (
	"math/big"
	"time"
)

type TransactionPayment struct {
	Id            *big.Int
	TransactionId uint64
	From          string
	To            string
	Index         uint16
	Amount        *big.Int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
