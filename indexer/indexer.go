package indexer

import (
	"context"
	"fmt"
	"go-bonotans/blockchain"
	"go-bonotans/model"
	"go-bonotans/repository"
	"math/big"
)

type Indexer interface {
	Index(startBlock *big.Int, numBlocks int)
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

func (indexer *DefaultIndexer) Index(startBlock *big.Int, numBlocks int) {
	ctx := context.Background()

	block := model.Block{
		Hash:       "H1",
		ParentHash: "PH1",
		Number:     big.NewInt(1),
	}

	b, err := indexer.blockRepository.Process(ctx, &block)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	t := block.Transactions[0]
	fmt.Printf("Chain %s", t.Chain())

	for _, transaction := range block.Transactions {
		indexer.transactionRepository.Save(ctx, transaction)

		for _, payment := range transaction.Payments {
			indexer.transactionPaymentRepository.Save(ctx, payment)
		}
	}

	fmt.Printf("SUCCESS: inserted block id: %s", b.Id)
}
