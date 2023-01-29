package indexer

import (
	"context"
	"fmt"
	"go-bonotans/blockchain"
	"go-bonotans/repository"
	"math/big"
)

type Indexer interface {
	Index(startBlock int64, numBlocks int)
}

type DefaultIndexer struct {
	blockchain                   blockchain.Blockchain
	blockRepository              *repository.BlockRepository
	transactionRepository        *repository.TransactionRepository
	transactionPaymentRepository *repository.TransactionPaymentRepository
	blockPositionRepository      *repository.IndexerPositionRepository
}

func NewDefaultIndexer(
	blockchain blockchain.Blockchain,
	blockRepository *repository.BlockRepository,
	transactionRepository *repository.TransactionRepository,
	transactionPaymentRepository *repository.TransactionPaymentRepository) Indexer {
	return &DefaultIndexer{
		blockchain:                   blockchain,
		blockRepository:              blockRepository,
		transactionRepository:        transactionRepository,
		transactionPaymentRepository: transactionPaymentRepository,
	}
}

func (indexer *DefaultIndexer) Index(startBlock int64, numBlocks int) {
	ctx := context.Background()

	for i := 0; i < numBlocks; i++ {
		blockNum := startBlock + int64(i)
		block := indexer.blockchain.GetBlock(ctx, big.NewInt(blockNum))
		fmt.Println(block.Hash)

		_, err := indexer.blockRepository.Process(ctx, block)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		for _, transaction := range block.Transactions {
			savedTransaction, err := indexer.transactionRepository.Save(ctx, transaction)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}

			for _, payment := range transaction.Payments {
				payment.TransactionId = savedTransaction.Id
				_, err := indexer.transactionPaymentRepository.Save(ctx, payment)
				if err != nil {
					fmt.Printf("Error: %s", err)
					return
				}
			}
		}
	}

}
