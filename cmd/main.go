package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-bonotans/di"
	//_ "go-bonotans/scanner"
	"math/big"
)

func main() {
	idx, err := di.InitializeIndexer()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	idx.Index(big.NewInt(0), 100)
}
