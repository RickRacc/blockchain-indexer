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
	indexerPositionRepository    *repository.IndexerPositionRepository
	sequencerPositionRepository  *repository.SequencerPositionRepository
}

func NewDefaultIndexer(
	blockchain blockchain.Blockchain,
	blockRepository *repository.BlockRepository,
	indexerPositionRepository *repository.IndexerPositionRepository,
	sequencerPositionRepository *repository.SequencerPositionRepository,
	transactionRepository *repository.TransactionRepository,
	transactionPaymentRepository *repository.TransactionPaymentRepository,
) Indexer {
	return &DefaultIndexer{
		blockchain:                   blockchain,
		blockRepository:              blockRepository,
		indexerPositionRepository:    indexerPositionRepository,
		sequencerPositionRepository:  sequencerPositionRepository,
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

func (indexer *DefaultIndexer) sequenceBlocks() error {
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
	ctx := context.Background()
	sequencerPosition, err := indexer.sequencerPositionRepository.GetCurrentPosition(ctx, coin.ETH)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	var currentPosition int64
	var block, nextBlock *model.Block
	if sequencerPosition == nil {
		// No sequencer position indexed, find the first block from db
		block, err = indexer.blockRepository.GetFirstBlock(ctx)
		if err != nil {
			return err
		}

		currentPosition = block.Number
		sequencerPosition := &model.SequencerPosition{
			CoinType: coin.ETH,
			Position: currentPosition,
		}
		indexer.sequencerPositionRepository.SaveCurrentPosition(ctx, sequencerPosition)
	}

	nextBlock, err = indexer.blockRepository.GetBlock(ctx, currentPosition)
	if err != nil {
		return err
	}

	if nextBlock.ParentHash != block.Hash {
		// fork condition
		position := currentPosition
		networkBlock := indexer.blockchain.GetBlock(ctx, big.NewInt(position))

		for networkBlock.Hash != block.Hash && networkBlock.ParentHash != block.ParentHash {
			position = position - 1
			block, err = indexer.blockRepository.GetBlock(ctx, position)
			if err != nil {

			}
			if block == nil {
				// need to start from beginning
			}
			networkBlock = indexer.blockchain.GetBlock(ctx, big.NewInt(position))
		}

		var lastBlock *model.Block
		lastBlock, err = indexer.blockRepository.GetLastBlock(ctx)

		for i := position + 1; i < lastBlock.Number; i++ {
			indexer.blockRepository.Delete(ctx, i)
		}

		// initiate indexing and sequencing from position + 1
	}

	//if (sequencerPosition.Position)
	return nil
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
	_, err = indexer.indexerPositionRepository.SaveCurrentPosition(ctx, position)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	blockChan <- blockNum
}
