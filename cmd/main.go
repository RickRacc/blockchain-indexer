package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-bonotans/di"
)

func main() {
	idx, err := di.InitializeIndexer()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	startBlock := int64(600004)
	numBlocks := 1

	idx.Index(startBlock, numBlocks)
}
