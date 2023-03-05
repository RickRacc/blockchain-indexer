package model

import "time"

type SequencerPosition struct {
	Id        uint64
	CoinType  int16
	Position  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
