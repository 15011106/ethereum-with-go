package main

import (
	"ethereum_with_go/chapter1"
	"fmt"
)

func main() {

	// Get Tx
	err, tx := chapter1.GetTx()
	if err != nil {
		panic(err)
	}

	fmt.Println(tx)

}
