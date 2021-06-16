package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMnemonic_24(t *testing.T) {
	a := assert.New(t)

	mnemonic, err := generateMnemonic(false)
	a.Nil(err)
	a.Len(strings.Split(mnemonic, " "), 24)
}

func TestGenerateMnemonic_12(t *testing.T) {
	a := assert.New(t)

	mnemonic, err := generateMnemonic(true)
	a.Nil(err)
	a.Len(strings.Split(mnemonic, " "), 12)
}
