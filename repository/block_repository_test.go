package repository

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"go-bonotans/di"
	"go-bonotans/model"
	"math/big"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var repo *BlockRepository
var pool *sql.DB

func setup() error {
	var err error
	pool, err = di.ProvideDbPool()
	if err != nil {
		return err
	}

	repo = di.ProvideBlockRepository(pool)
	if err != nil {
		return err
	}

	return nil
}

func shutdown() {

}

func TestCreateBlock(t *testing.T) {
	assert := assert.New(t)
	block := model.Block{
		ParentHash:   "parenthash",
		Hash:         "hash",
		Number:       new(big.Int).SetInt64(1),
		Transactions: nil,
	}

	b, err := repo.Process(context.Background(), &block)
	assert.NoError(err, "Saving block returned an error")
	assert.NotNil(b.Id)
	assert.NotNil(b.CreatedAt)
	assert.NotNil(b.UpdatedAt)
}
