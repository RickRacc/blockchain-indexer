package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/test"
	"testing"
)

type TransactionPaymentRepositoryTestSuite struct {
	test.BaseTestSuite
	blockRepo              *BlockRepository
	transactionRepo        *TransactionRepository
	transactionPaymentRepo *TransactionPaymentRepository
}

func (suite *TransactionPaymentRepositoryTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.blockRepo = NewBlockRepository(suite.Pool)
	suite.transactionRepo = NewTransactionRepository(suite.Pool)
	suite.transactionPaymentRepo = NewTransactionPaymentRepository(suite.Pool)
}

func (suite *TransactionPaymentRepositoryTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}

func (suite *TransactionPaymentRepositoryTestSuite) TestCreateTransaction() {
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
	if err != nil {
		panic(err)
	}

	transactionPayment := test.GetTransactionPayment()
	transactionPayment.TransactionId = transaction.Id
	transactionPayment, err = suite.transactionPaymentRepo.Save(ctx, transactionPayment)

	assert.NoError(err)
	assert.NotNil(transactionPayment.Id)
	assert.NotNil(transactionPayment.CreatedAt)
	assert.NotNil(transactionPayment.UpdatedAt)
}

func TestTransactionPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionPaymentRepositoryTestSuite))
}
