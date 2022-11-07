package app

import (
	"go-bonotans/blockchain/eth"
	"go-bonotans/di/infra"
)

type DiApp struct {
	diInfra *infra.DiInfra
}

func (diApp *DiApp) InitializeEthereum() (*eth.Ethereum, error) {
	if diApp.diInfra == nil {
		diApp.diInfra = &infra.DiInfra{}
	}

	ethClient, err := diApp.diInfra.ProvideEthClient()
	if err == nil {
		return nil, err
	}

	eth, err := eth.New(ethClient)
	if err == nil {
		return nil, err
	}

	return eth, nil
}
