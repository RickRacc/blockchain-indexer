package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/model"
	"go-bonotans/test"
	"math/big"
	"testing"
)

type BlockRepositoryTestSuite struct {
	test.BaseTestSuite
	repo *BlockRepository
}

func (suite *BlockRepositoryTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.repo = NewBlockRepository(suite.Pool)
}

func (suite *BlockRepositoryTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}

func (suite *BlockRepositoryTestSuite) TestCreateBlock() {
	assert := assert.New(suite.T())
	block := model.Block{
		ParentHash:   "parenthash",
		Hash:         "hash",
		Number:       new(big.Int).SetInt64(1),
		Transactions: nil,
	}

	b, err := suite.repo.Process(context.Background(), &block)
	assert.NoError(err, "Saving block returned an error")
	assert.NotNil(b.Id)
	assert.NotNil(b.CreatedAt)
	assert.NotNil(b.UpdatedAt)
}

func TestBlockRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BlockRepositoryTestSuite))
}
