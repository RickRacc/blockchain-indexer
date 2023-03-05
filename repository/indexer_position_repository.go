package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
)

type IndexerPositionRepository struct {
	pool *sql.DB
}

func NewIndexerPositionRepository(pool *sql.DB) *IndexerPositionRepository {
	return &IndexerPositionRepository{
		pool: pool,
	}
}

func (repo *IndexerPositionRepository) GetCurrentPosition(ctx context.Context, coinType int16) (*model.IndexerPosition, error) {
	stmt := fmt.Sprintf("select %s from indexer_position where coin_type=%v",
		INDEXER_POSITION_SELECT_COLS, coinType)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *IndexerPositionRepository) SaveCurrentPosition(ctx context.Context, indexerPosition *model.IndexerPosition) (*model.IndexerPosition, error) {
	// TODo: Update if the entry already exists
	stmt := fmt.Sprintf("insert into indexer_position (%s) values ('%v', '%v') returning %s",
		INDEXER_POSITION_INSERT_COLS, indexerPosition.CoinType, indexerPosition.Position, INDEXER_POSITION_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *IndexerPositionRepository) Read(row *sql.Row) (*model.IndexerPosition, error) {
	var indexerPosition model.IndexerPosition

	err := row.Scan(&indexerPosition.Id, &indexerPosition.CoinType, &indexerPosition.Position, &indexerPosition.CreatedAt, &indexerPosition.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &indexerPosition, nil
}
