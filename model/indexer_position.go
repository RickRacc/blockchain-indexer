package model

import (
	"time"
)

type IndexerPosition struct {
	Id        uint64
	CoinType  int16
	Position  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
