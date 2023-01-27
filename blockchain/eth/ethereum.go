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

func New(client *ethclient.Client) (*Ethereum, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	return &Ethereum{
		Client:  client,
		ChainID: chainID,
	}, nil
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
		Number:     blockNumber.Int64(),
		//Time:         b.Time(),
		Transactions: nil,
	}

	blockTxns := b.Transactions()
	transactions := make([]*model.EthTransaction, blockTxns.Len())
	for i, txn := range blockTxns {
		fmt.Println("Processing", i)
		receipt, _ := eth.Client.TransactionReceipt(context.Background(), txn.Hash())
		transaction := model.EthTransaction{
			BaseTransaction: model.BaseTransaction{
				Hash: txn.Hash().Hex(),
				Fee:  new(big.Int).Mul(new(big.Int).SetUint64(txn.Gas()), txn.GasPrice()),
			},
			Gas:                new(big.Int).SetUint64(txn.Gas()),
			GasPrice:           txn.GasPrice(),
			IsContractCreation: false,
		}

		payment := model.TransactionPayment{
			Amount: txn.Value(),
		}
		if txn.To() != nil {
			payment.To = txn.To().Hex()
		} else {
			// This is a contract creation call and the address is available in receipt
			payment.To = receipt.ContractAddress.Hex()
			transaction.IsContractCreation = true
		}

		if msg, err := txn.AsMessage(types.NewEIP155Signer(eth.ChainID), nil); err == nil {
			payment.From = msg.From().Hex()
		}
		transaction.Payments = []*model.TransactionPayment{&payment}
		transactions[i] = &transaction
	}

	block.Transactions = transactions

	return &block
}
