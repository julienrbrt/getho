package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateWallet(t *testing.T) {
	a := assert.New(t)

	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := generateWallet(mnemonic)
	a.Nil(err)
	a.Equal(wallet, "0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947")
}
