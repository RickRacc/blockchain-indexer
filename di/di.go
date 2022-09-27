//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
	"go-bonotans/blockchain/eth"
	"go-bonotans/config"
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

func InitializeEthereum() *eth.Ethereum {
	wire.Build(eth.New, provideEthClient, ethRpcUrl)
	return &eth.Ethereum{}
}
