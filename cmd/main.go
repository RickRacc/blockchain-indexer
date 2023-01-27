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

	idx.Index(1, 100)
}
