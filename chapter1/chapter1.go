package chapter1

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

func GetTx() (err error, tx *types.Transaction) {

	// Ethereum RPC endpoint URL (ankr Rpc of mainnet)
	url := "https://rpc.ankr.com/eth/"

	// Create RPC client
	var c *rpc.Client
	c, err = rpc.DialContext(context.TODO(), url)
	if err != nil {

		return errors.Wrap(err, "rpc Client error"), nil
	}

	// Create Ethereum client
	client := ethclient.NewClient(c)

	var pending bool
	tx, pending, err = client.TransactionByHash(context.TODO(), common.HexToHash("0xb822fa2ab0f5bb78f5912fff954a05b15b39b877393fb2e8bd2b8de727c18982"))

	if err != nil {
		return errors.Wrap(err, "Transaction error"), nil
	}

	if pending {
		return err, nil
	}

	return nil, tx

}
