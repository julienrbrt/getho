package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	numberWords int
)

func init() {
	flag.IntVar(&numberWords, "n", 24, "mnemonic number of words (default 24)")
	flag.Parse()
}

func main() {
	// parse arguments
	var shortSize bool
	switch numberWords {
	case 12:
		shortSize = true
	case 24:
		shortSize = false
	default:
		log.Fatalf("%d word mnemonic is not supported: use 12 or 24", numberWords)
		return
	}

	ethClient, err := connectETHNode("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	for {
		mnemonic, err := generateMnemonic(shortSize)
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

			if err := saveResult(address.Hex(), mnemonic, balance); err != nil {
				os.Exit(0)
			}
		}

		fmt.Print(".")
	}
}

func saveResult(address, mnemonic string, balance int64) error {
	filename := "result.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed opening %s: %w", filename, err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("| %s | %d | %s |\n", address, balance, mnemonic)); err != nil {
		return fmt.Errorf("failed appeding to %s: %w", filename, err)
	}

	return nil
}
