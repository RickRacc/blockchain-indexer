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
	//tx := ctx.Value(TRANSACTION)
	//var err error
	//if tx == nil {
	//	tx, err = repo.pool.Begin()
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	//var ids [2]int64
	//var id big.Int
	//sql := fmt.Sprintf("insert into block (\"Hash\", \"ParentHash\", \"Number\") values ('%s', '%s', %s) returning *", //"Number"
	//	block.Hash, block.ParentHash, block.Number)

	sql := fmt.Sprintf("insert into block (%s) values ('%s', '%s', %s) returning %s",
		BLOCK_INSERT_COLS, block.Hash, block.ParentHash, block.Number, BLOCK_SELECT_COLS)

	//sql := "insert into block (\"Hash\", \"ParentHash\", \"Number\") values ('a', 'aa', 11), ('b', 'bb', 22) returning \"Id\""
	//block.Id = new(big.Int)
	//var b model.Block
	//var num []byte
	//err := repo.pool.QueryRowContext(ctx, sql).Scan(&num)

	row := repo.pool.QueryRowContext(ctx, sql)

	b, err := repo.Read(row)
	if err != nil {
		return nil, err
	}

	// Write transaction and transaction payments

	//block.Id = new(big.Int).SetUint64(id)
	//fmt.Printf("%v", ret)
	return b, nil
	//counter := 0
	//for ret.Next() {
	//	var id int64
	//	ret.Scan(&id)
	//	ids[counter] = id
	//	counter += 1
	//}

	//tx.Commit()
	//id, err := result.LastInsertId()
	//if err != nil {
	//	return nil, err
	//}

	//return big.NewInt(ids[0]), nil
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
