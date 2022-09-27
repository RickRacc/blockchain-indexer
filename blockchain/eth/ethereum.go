package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-bonotans/model"
	"math/big"
)

type Ethereum struct {
	Client  *ethclient.Client
	ChainID *big.Int
}

func New(client *ethclient.Client) *Ethereum {
	chainID, _ := client.NetworkID(context.Background())
	return &Ethereum{
		Client:  client,
		ChainID: chainID,
	}
}

func (eth *Ethereum) GetTip(ctx context.Context) *big.Int {
	number, _ := eth.Client.BlockNumber(ctx)
	return big.NewInt(int64(number))
}

func (eth *Ethereum) GetBlock(ctx context.Context, blockNumber *big.Int) *model.Block {
	b, _ := eth.Client.BlockByNumber(ctx, blockNumber)
	block := model.Block{
		ParentHash: b.ParentHash().Hex(),
		Hash:       b.Hash().Hex(),
		Number:     blockNumber,
		//Time:         b.Time(),
		Transactions: nil,
	}

	blockTxns := b.Transactions()
	transactions := make([]*model.EthTransaction, blockTxns.Len())
	for i, txn := range blockTxns {
		fmt.Println("Processing", i)
		receipt, _ := eth.Client.TransactionReceipt(context.Background(), txn.Hash())
		//fmt.Println(receipt)
		transaction := model.EthTransaction{
			Transaction: model.Transaction{
				ID:     txn.Hash().Hex(),
				Amount: txn.Value(),
				Fee:    new(big.Int).Mul(new(big.Int).SetUint64(txn.Gas()), txn.GasPrice()),
			},
			Gas:                txn.Gas(),
			GasPrice:           txn.GasPrice(),
			IsContractCreation: false,
		}

		if txn.To() != nil {
			transaction.To = txn.To().Hex()
		} else {
			// This is a contract creation call and the address is available in receipt
			transaction.To = receipt.ContractAddress.Hex()
			transaction.IsContractCreation = true
		}

		if msg, err := txn.AsMessage(types.NewEIP155Signer(eth.ChainID), nil); err == nil {
			transaction.From = msg.From().Hex()
		}
		transactions[i] = &transaction
	}

	block.Transactions = transactions

	return &block
}
