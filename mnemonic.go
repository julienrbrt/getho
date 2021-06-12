package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

// generateMnemonic for memorization or user-friendly seeds
func generateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", fmt.Errorf("error while creating entropy: %w", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("error while generating mnemonic: %w", err)
	}

	return mnemonic, nil
}
