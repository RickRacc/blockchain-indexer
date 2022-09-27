package main

import (
	"context"
	"fmt"
	"go-bonotans/di"
)

func main() {
	ethIndexer := di.InitializeEthereum()
	ctx := context.Background()
	blockNumber := ethIndexer.GetTip(ctx)
	fmt.Println("Tip:", blockNumber)
	block := ethIndexer.GetBlock(ctx, blockNumber)
	fmt.Println("Block:", block)
}
