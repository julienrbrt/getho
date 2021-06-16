package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethClient struct {
	client *ethclient.Client
}

func connectETHNode(endpoint string) (*ethClient, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connecto to eth node %s: %w", endpoint, err)
	}
	log.Println("successfully connected to ethereum node 🎉")

	return &ethClient{client}, nil
}

func (e *ethClient) checkBalance(ctx context.Context, address common.Address) (int64, error) {
	balance, err := e.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return 0, fmt.Errorf("error when checking balance of %s: %w", address.Hex(), err)
	}

	return balance.Int64(), nil
}
