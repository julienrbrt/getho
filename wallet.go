package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func generateWallet(mnemonic string) (*common.Address, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("error while generating wallet from mnemonic: %s", err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing devivation path: %s", err)
	}

	return &account.Address, nil
}
