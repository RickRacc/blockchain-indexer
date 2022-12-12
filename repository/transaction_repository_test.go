package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/test"
	"testing"
)

type TransactionRepositoryTestSuite struct {
	test.BaseTestSuite
	blockRepo       *BlockRepository
	transactionRepo *TransactionRepository
}

func (suite *TransactionRepositoryTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.blockRepo = NewBlockRepository(suite.Pool)
	suite.transactionRepo = NewTransactionRepository(suite.Pool)
}

func (suite *TransactionRepositoryTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}

func (suite *TransactionRepositoryTestSuite) TestCreateTransaction() {
	assert := assert.New(suite.T())
	block := test.GetBlock()
	ctx := context.Background()
	b, err := suite.blockRepo.Process(ctx, block)
	if err != nil {
		panic(err)
	}

	transaction := test.GetEthTransaction()
	transaction.BlockNumber = b.Number
	transaction, err = suite.transactionRepo.Save(ctx, transaction)
	assert.NoError(err)
	assert.NotNil(transaction.Id)
	assert.NotNil(transaction.CreatedAt)
	assert.NotNil(transaction.UpdatedAt)
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
