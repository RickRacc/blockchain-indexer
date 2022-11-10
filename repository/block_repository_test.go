package repository

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-bonotans/model"
	"go-bonotans/test"
	"math/big"
	"testing"
)

type BlockTestSuite struct {
	test.BaseTestSuite
	repo *BlockRepository
	pool *sql.DB
}

//func (suite *BlockTestSuite) SetupAllSuite() {
//	var err error
//
//	diInfra := infra.DiInfra{}
//	suite.pool, err = diInfra.ProvideDbPool()
//	if err != nil {
//		os.Exit(1)
//	}
//
//	suite.repo = NewBlockRepository(suite.pool)
//}
//
//func (suite *BlockTestSuite) TearDownAllSuite() {
//	err := suite.pool.Close()
//	if err != nil {
//		os.Exit(1)
//	}
//}

func (suite *BlockTestSuite) SetupAllSuiteInternal() {
	suite.repo = NewBlockRepository(suite.pool)
}

func (suite *BlockTestSuite) TearDownAllSuiteInternal() {
}

func (suite *BlockTestSuite) TestCreateBlock() {
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

func TestBlockTestSuite(t *testing.T) {
	suite.Run(t, new(BlockTestSuite))
}
