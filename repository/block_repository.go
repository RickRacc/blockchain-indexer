package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
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
	stmt := fmt.Sprintf("insert into block (%s) values ('%s', '%s', %v) returning %s",
		BLOCK_INSERT_COLS, block.Hash, block.ParentHash, block.Number, BLOCK_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) Delete(ctx context.Context, blockNumber int64) (*model.Block, error) {
	stmt := fmt.Sprintf("delete from block where block_number=%v returning %s",
		blockNumber, BLOCK_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) Read(row *sql.Row) (*model.Block, error) {
	var block model.Block

	err := row.Scan(&block.Id, &block.Hash, &block.ParentHash, &block.Number, &block.CreatedAt, &block.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
