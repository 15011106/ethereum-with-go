package chapter2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func DeploySmartContract() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal("Error connecting to the Ethereum client:", err)
	}

	// Private key (ensure this key has enough balance)
	privateKey, err := crypto.HexToECDSA("...")
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Check account balance
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal("Error getting balance:", err)
	}
	fmt.Printf("Account balance: %s\n", balance.String())

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Error getting nonce:", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error suggesting gas price:", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Error getting chain ID:", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Error creating transaction opts:", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Amount of ETH to send with the transaction
	auth.GasLimit = uint64(5000000) // Reasonable gas limit
	auth.GasPrice = gasPrice

	// Adjust contract deployment function as necessary
	address, tx, _, err := DeployChapter2(auth, client)
	if err != nil {
		log.Fatal("Error deploying contract:", err)
	}

	fmt.Println("Contract deployed to address:", address.Hex())
	fmt.Println("Transaction hash:", tx.Hash().Hex())
}
