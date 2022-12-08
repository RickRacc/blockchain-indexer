package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/model"
	"go-bonotans/test"
	"testing"
)

type IndexerPositionRepositoryTestSuite struct {
	test.BaseTestSuite
	repo *IndexerPositionRepository
}

func (suite *IndexerPositionRepositoryTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.repo = NewIndexerPositionRepository(suite.Pool)
}

func (suite *IndexerPositionRepositoryTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}

func (suite *IndexerPositionRepositoryTestSuite) TestSaveCurrentPosition() {
	assert := assert.New(suite.T())
	indexerPosition := model.IndexerPosition{
		CoinType: 0,
		Position: 1,
	}

	position, err := suite.repo.SaveCurrentPosition(context.Background(), &indexerPosition)
	assert.NoError(err, "Saving indexer position returned an error")
	assert.NotNil(position.Id)
	assert.NotNil(position.CreatedAt)
	assert.NotNil(position.UpdatedAt)
}

func TestIndexerPositionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(IndexerPositionRepositoryTestSuite))
}
