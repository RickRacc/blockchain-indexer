package infra

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-bonotans/config"
)

type DiInfra struct {
	pool      *sql.DB
	ethClient *ethclient.Client
}

func (infra *DiInfra) ProvideDbPool() (*sql.DB, error) {
	if infra.pool == nil {
		cfg := config.Config()
		host := cfg.String("db.host") // localhost
		port := cfg.Int("db.port")    // 5432
		user := cfg.String("db.user") //"postgres"
		password := cfg.String("db.password")
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

		infra.pool = pool
	}

	return infra.pool, nil
}

func EthRpcUrl() string {
	return config.Config().String("eth.rpcUrl")
}

func (infra *DiInfra) ProvideEthClient() (*ethclient.Client, error) {
	if infra.ethClient == nil {
		rpcUrl := config.Config().String("eth.rpcUrl")
		ethClient, err := ethclient.Dial(rpcUrl)
		if err != nil {
			return nil, err
		}

		infra.ethClient = ethClient
	}

	return infra.ethClient, nil
}
