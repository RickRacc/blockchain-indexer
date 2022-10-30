//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
	"go-bonotans/blockchain/eth"
	"go-bonotans/config"
	"go-bonotans/indexer"
	"go-bonotans/repository"
	"log"
)

//ethConfig

func ethRpcUrl() string {
	return config.Config().String("eth.rpcUrl")
}

func provideEthClient(rpcUrl string) *ethclient.Client {
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	return ethClient
}

func provideDbPool() (*sql.DB, error) {
	cfg := config.Config()
	host := cfg.String("db.host")         // localhost
	port := cfg.Int("db.port")            // 5432
	user := cfg.String("db.user")         //"postgres"
	password := cfg.String("db.password") //"Test@12344"
	sslmode := cfg.String("db.sslmode")
	dbname := cfg.String("db.name") //"bonotans"

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	pool, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	return pool, nil
}

func provideBlockRepository(pool *sql.DB) *repository.BlockRepository {
	return repository.NewBlockRepository(pool)
}

func provideTransactionRepository(pool *sql.DB) *repository.TransactionRepository {
	return repository.NewTransactionRepository(pool)
}

func provideTransactionPaymentRepository(pool *sql.DB) *repository.TransactionPaymentRepository {
	return repository.NewTransactionPaymentRepository(pool)
}

func InitializeEthereum() *eth.Ethereum { //
	wire.Build(eth.New, provideEthClient, ethRpcUrl)
	return &eth.Ethereum{}
}

func InitializeIndexer() (indexer.Indexer, error) {
	//wire.Build(indexer.NewDefaultIndexer, InitializeEthereum, provideBlockRepository,
	//	provideTransactionRepository, provideTransactionPaymentRepository, provideDbPool)
	//
	//return &indexer.DefaultIndexer{}

	pool, err := provideDbPool()
	if err != nil {
		return nil, err
	}
	return indexer.NewDefaultIndexer(
		InitializeEthereum(),
		provideBlockRepository(pool),
		provideTransactionRepository(pool),
		provideTransactionPaymentRepository(pool),
	), nil
}
