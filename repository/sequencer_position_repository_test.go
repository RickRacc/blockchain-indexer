package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/coin"
	"go-bonotans/model"
	"go-bonotans/test"
	"testing"
)

type SequencerPositionRepositoryTestSuite struct {
	test.BaseTestSuite
	repo *SequencerPositionRepository
}

func (suite *SequencerPositionRepositoryTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.repo = NewSequencerPositionRepository(suite.Pool)
}

func (suite *SequencerPositionRepositoryTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}

func (suite *SequencerPositionRepositoryTestSuite) TestSaveCurrentPosition() {
	assert := assert.New(suite.T())
	sequencerPosition := model.SequencerPosition{
		CoinType: coin.ETH,
		Position: 1,
	}

	position, err := suite.repo.SaveCurrentPosition(context.Background(), &sequencerPosition)
	assert.NoError(err, "Saving sequencer position returned an error")
	assert.NotNil(position.Id)
	assert.Equal(position.CoinType, sequencerPosition.CoinType)
	assert.Equal(position.Position, sequencerPosition.Position)
	assert.NotNil(position.CreatedAt)
	assert.NotNil(position.UpdatedAt)
}

func (suite *SequencerPositionRepositoryTestSuite) TestGetCurrentPosition() {
	assert := assert.New(suite.T())
	const coinType int16 = coin.ETH
	sequencerPosition := model.SequencerPosition{
		CoinType: coinType,
		Position: 1,
	}

	ctx := context.Background()
	_, err := suite.repo.SaveCurrentPosition(ctx, &sequencerPosition)
	assert.NoError(err, "Saving sequencer position returned an error")

	var position *model.SequencerPosition
	position, err = suite.repo.GetCurrentPosition(ctx, coinType)
	assert.NoError(err, "Getting sequencer current position returned an error")
	assert.NotNil(position.Id)
	assert.Equal(position.CoinType, sequencerPosition.CoinType)
	assert.Equal(position.Position, sequencerPosition.Position)
	assert.NotNil(position.CreatedAt)
	assert.NotNil(position.UpdatedAt)
}

func TestSequencerPositionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SequencerPositionRepositoryTestSuite))
}
