package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-bonotans/model"
)

type SequencerPositionRepository struct {
	pool *sql.DB
}

func NewSequencerPositionRepository(pool *sql.DB) *SequencerPositionRepository {
	return &SequencerPositionRepository{
		pool: pool,
	}
}

func (repo *SequencerPositionRepository) GetCurrentPosition(ctx context.Context, coinType int16) (*model.SequencerPosition, error) {
	stmt := fmt.Sprintf("select %s from sequencer_position where coin_type=%v",
		SEQUENCER_POSITION_SELECT_COLS, coinType)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *SequencerPositionRepository) SaveCurrentPosition(ctx context.Context, sequencerPosition *model.SequencerPosition) (*model.SequencerPosition, error) {
	// TODo: Update if the entry already exists
	stmt := fmt.Sprintf("insert into sequencer_position (%s) values ('%v', '%v') returning %s",
		SEQUENCER_POSITION_INSERT_COLS, sequencerPosition.CoinType, sequencerPosition.Position, SEQUENCER_POSITION_SELECT_COLS)

	row := repo.pool.QueryRowContext(ctx, stmt)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (repo *SequencerPositionRepository) Read(row *sql.Row) (*model.SequencerPosition, error) {
	var sequencerPosition model.SequencerPosition

	err := row.Scan(&sequencerPosition.Id, &sequencerPosition.CoinType, &sequencerPosition.Position, &sequencerPosition.CreatedAt, &sequencerPosition.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &sequencerPosition, nil
}
