package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
	"math/big"
)

type BlockRepository struct {
	pool *sql.DB
}

func NewBlockRepository(pool *sql.DB) *BlockRepository {
	return &BlockRepository{
		pool: pool,
	}
}

func (repo *BlockRepository) Process(ctx context.Context, block *model.Block) (*model.Block, error) {
	stmt := fmt.Sprintf("insert into block (%s) values ('%s', '%s', %s) returning %s",
		BLOCK_INSERT_COLS, block.Hash, block.ParentHash, block.Number, BLOCK_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) Read(row *sql.Row) (*model.Block, error) {
	var number []byte
	var block model.Block

	err := row.Scan(&block.Id, &block.Hash, &block.ParentHash, &number, &block.CreatedAt, &block.UpdatedAt)
	if err != nil {
		return nil, err
	}
	block.Number = new(big.Int).SetBytes(number)

	return &block, nil
}
