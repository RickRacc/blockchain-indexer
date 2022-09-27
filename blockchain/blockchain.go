package blockchain

import (
	"context"
	"go-bonotans/model"
	"math/big"
)

type Blockchain interface {
	GetTip(ctx context.Context) *big.Int
	GetBlock(ctx context.Context, blockNumber *big.Int) *model.Block
}
