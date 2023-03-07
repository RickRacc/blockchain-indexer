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
	repo.Delete(ctx, block.Number)
	stmt := fmt.Sprintf("insert into block (%s) values ('%s', '%s', %v) returning %s",
		BLOCK_INSERT_COLS, block.Hash, block.ParentHash, block.Number, BLOCK_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) GetFirstBlock(ctx context.Context) (*model.Block, error) {
	stmt := fmt.Sprintf("select (%s) from block where number=(select min(number) from block)", BLOCK_SELECT_COLS)
	row := repo.pool.QueryRowContext(ctx, stmt)
	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) GetLastBlock(ctx context.Context) (*model.Block, error) {
	stmt := fmt.Sprintf("select (%s) from block where number=(select max(number) from block)", BLOCK_SELECT_COLS)
	row := repo.pool.QueryRowContext(ctx, stmt)
	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) GetBlock(ctx context.Context, number int64) (*model.Block, error) {
	stmt := fmt.Sprintf("select (%s) from block where number=%v", BLOCK_SELECT_COLS, number)
	row := repo.pool.QueryRowContext(ctx, stmt)
	b, err := repo.Read(row)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return b, nil
}

func (repo *BlockRepository) Delete(ctx context.Context, blockNumber int64) (*model.Block, error) {
	stmt := fmt.Sprintf("delete from block where number=%v returning %s",
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
