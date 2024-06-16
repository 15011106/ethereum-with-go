package main

import (
	"ethereum_with_go/chapter2"
)

func main() {

	//// Get Tx
	//err, tx := chapter1.GetTx()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%+v\n", tx)
	//
	//// Get from (
	//sender, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	//fmt.Println(sender)

	chapter2.DeploySmartContract()
}
