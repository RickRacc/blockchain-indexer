package repository

import "database/sql"

type BlockPositionRepository struct {
	pool *sql.DB
}

func NewBlockCounterRepository(pool *sql.DB) *BlockPositionRepository {
	return &BlockPositionRepository{
		pool: pool,
	}
}

func (repo *BlockPositionRepository) GetCurrentPosition() int64 {
	return 0
}

func (repo *BlockPositionRepository) SaveCurrentPosition(currentPosition int64) error {
	return nil
}
