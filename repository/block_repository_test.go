package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/test"
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
	block := test.GetBlock()

	b, err := suite.repo.Process(context.Background(), block)
	assert.NoError(err, "Saving block returned an error")
	assert.NotNil(b.Id)
	assert.NotNil(b.CreatedAt)
	assert.NotNil(b.UpdatedAt)
}

func (suite *BlockRepositoryTestSuite) TestGetFirstBlock() {
	assert := assert.New(suite.T())
	blocks := test.GetBlocks()

	ctx := context.Background()
	for _, block := range blocks {
		_, err := suite.repo.Process(ctx, &block)
		assert.NoError(err, fmt.Sprintf("Saving block returned an error: %s", err.Error()))
	}

	b, err := suite.repo.GetFirstBlock(ctx)
	assert.NoError(err, fmt.Sprintf("Saving block returned an error", err.Error()))
	assert.NotNil(b.Id)
	assert.Equal(1, b.Number)
	assert.NotNil(b.CreatedAt)
	assert.NotNil(b.UpdatedAt)
}

func TestBlockRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BlockRepositoryTestSuite))
}
