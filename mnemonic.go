package main

import (
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

// generateMnemonic for memorization or user-friendly seeds
func generateMnemonic(shortSize bool) (string, error) {
	bitSize := 256
	if shortSize {
		bitSize = 128
	}

	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", fmt.Errorf("error while creating entropy: %w", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("error while generating mnemonic: %w", err)
	}

	return mnemonic, nil
}
