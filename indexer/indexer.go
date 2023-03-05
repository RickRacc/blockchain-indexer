package indexer

import (
	"context"
	"fmt"
	"go-bonotans/blockchain"
	"go-bonotans/coin"
	"go-bonotans/model"
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
	blockPositionRepository *repository.IndexerPositionRepository,
	transactionRepository *repository.TransactionRepository,
	transactionPaymentRepository *repository.TransactionPaymentRepository) Indexer {
	return &DefaultIndexer{
		blockchain:                   blockchain,
		blockRepository:              blockRepository,
		blockPositionRepository:      blockPositionRepository,
		transactionRepository:        transactionRepository,
		transactionPaymentRepository: transactionPaymentRepository,
	}
}

func (indexer *DefaultIndexer) Index(startBlock int64, numBlocks int) {
	numRoutines := 100
	routineCounter := 0
	blockChans := make(chan int64, numRoutines)

	lastBlock := startBlock + int64(numBlocks)

	for i := startBlock; i < lastBlock; i++ {
		if routineCounter < numRoutines {
			go indexer.indexBlock(i, blockChans)
			i++
			routineCounter++
			continue
		}

		<-blockChans
		go indexer.indexBlock(i, blockChans)
	}

}

func (indexer *DefaultIndexer) sequenceBlocks() {
	/**
	 * last verified block
	fetch (block) and (block + 1)[nextBlock] from db
	if found and nextBlock.prevhash != block.prevhash
		fork condition
		fetch (block - 1) from both db and blockchain
		if (hash and prev hash not same, repeat until the hash is same, decrementing block)

		delete all blocks after the match
		re-fectch the blocks after the match
	*/
}

func (indexer *DefaultIndexer) indexBlock(blockNum int64, blockChan chan<- int64) {
	ctx := context.Background()

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

	position := &model.IndexerPosition{
		CoinType: coin.ETH,
		Position: blockNum,
	}
	indexer.blockPositionRepository.SaveCurrentPosition(ctx, position)

	blockChan <- blockNum
}
