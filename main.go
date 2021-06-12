package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// if len(os.Args) > 12 {
	// 	bip39.SetWordList(os.Args[1:])
	// }

	ethClient := ConnectEthereumNode("http://127.0.0.1:8545")

	for {
		mnemonic, err := generateMnemonic()
		if err != nil {
			log.Fatal(err)
		}

		address, err := generateWallet(mnemonic)
		if err != nil {
			log.Fatal(err)
		}

		balance, err := ethClient.checkBalance(context.Background(), *address)
		if err != nil {
			log.Printf("uh oh: %s\n", err)
		}

		if balance > 0 {
			log.Printf("🎉🎉🎉 address %s has %d balance 🎉🎉🎉", address.Hex(), balance)
			log.Printf("its private key is based on this mnemonic: '%s'", mnemonic)
			os.Exit(0)
		}

		fmt.Print(".")
	}
}

type ethClient struct {
	client *ethclient.Client
}

func ConnectEthereumNode(endpoint string) ethClient {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to ethereum node 🎉")

	return ethClient{client}
}

func (e *ethClient) checkBalance(ctx context.Context, address common.Address) (int64, error) {
	balance, err := e.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return 0, fmt.Errorf("error when checking balance of %s: %w", address.Hex(), err)
	}

	return balance.Int64(), nil
}
