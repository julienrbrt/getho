package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMnemonic(t *testing.T) {
	a := assert.New(t)

	mnemonic, err := generateMnemonic()
	a.Nil(err)
	a.Len(strings.Split(mnemonic, " "), 24)
}
